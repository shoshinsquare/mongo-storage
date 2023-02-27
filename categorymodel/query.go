package categorymodel

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3/operator"
	"github.com/ponlv/go-kit/mongodb"
	"github.com/ponlv/go-kit/mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create manga
func Create(ctx context.Context, col *Category) (interface{}, error) {
	nft := mongodb.Coll(col)
	id, err := nft.CreateWithCtx(ctx, col)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func FindAll(ctx context.Context) ([]*Category, error) {
	query := bson.D{
		{"$or", bson.A{
			bson.M{"is_delete": false},
			bson.M{"is_delete": nil},
		}},
	}

	sort := bson.M{
		"updated_at": -1,
	}

	pipeline := bson.A{}
	pipeline = utils.BsonAggregate(pipeline, operator.Match, query)
	pipeline = utils.BsonAggregate(pipeline, operator.Sort, sort)

	return findWithAggregate(ctx, pipeline)
}

func FindAllByListId(ctx context.Context, listId []string) ([]*Category, error) {
	var listObjectId []primitive.ObjectID
	for _, id := range listId {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		listObjectId = append(listObjectId, objID)
	}
	query := bson.D{
		{"$or", bson.A{
			bson.M{"is_delete": false},
			bson.M{"is_delete": nil},
		}},
		{"_id", bson.M{"$in": listObjectId}},
	}

	sort := bson.M{
		"updated_at": -1,
	}

	pipeline := bson.A{}
	pipeline = utils.BsonAggregate(pipeline, operator.Match, query)
	pipeline = utils.BsonAggregate(pipeline, operator.Sort, sort)

	return findWithAggregate(ctx, pipeline)
}
func FindAllPriority(ctx context.Context, limit, offset int64) ([]*Category, error) {
	query := bson.D{
		{"$or", bson.A{
			bson.M{"is_delete": false},
			bson.M{"is_delete": nil},
		}},
		{"priority", bson.M{"$gt": 0}},
	}

	sort := bson.M{
		"priority": 1,
	}

	pipeline := bson.A{}
	pipeline = utils.BsonAggregate(pipeline, operator.Match, query)
	pipeline = utils.BsonAggregate(pipeline, operator.Sort, sort)
	pipeline = utils.BsonAggregate(pipeline, operator.Skip, offset)
	pipeline = utils.BsonAggregate(pipeline, operator.Limit, limit)

	return findWithAggregate(ctx, pipeline)
}

func UpdateById(ctx context.Context, id string, update bson.D, opts ...*options.UpdateOptions) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	update = utils.BsonSet(update, "updated_at", time.Now().Unix())
	return updateOne(ctx, filter, update, opts...)
}

func DeleteById(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}

	update := bson.D{}
	update = utils.BsonSet(update, "is_delete", true)
	update = utils.BsonSet(update, "updated_at", time.Now().Unix())
	update = utils.BsonSet(update, "deleted_at", time.Now().Unix())

	return updateOne(ctx, filter, update)
}

// FindById find nft by token id
func FindById(ctx context.Context, id string) (*Category, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	col := mongodb.Coll(&Category{})

	nft := &Category{}
	err = col.FindOne(ctx, bson.M{"_id": objID}).Decode(nft)
	if err != nil {
		return nil, err
	}
	return nft, nil
}

// FindBySlug find nft by token id
func FindBySlug(ctx context.Context, slug string) (*Category, error) {
	col := mongodb.Coll(&Category{})

	nft := &Category{}
	err := col.FindOne(ctx, bson.M{"slug": slug}).Decode(nft)
	if err != nil {
		return nil, err
	}
	return nft, nil
}

func findWithAggregate(ctx context.Context, pipeline bson.A) ([]*Category, error) {
	col := mongodb.Coll(&Category{})
	var nfts []*Category
	cursor, err := col.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var u *Category
		if err = cursor.Decode(&u); err != nil {
			return nil, err
		} else {
			nfts = append(nfts, u)
		}
	}

	return nfts, nil
}

func updateOne(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) error {
	col := mongodb.Coll(&Category{})
	_, err := col.UpdateOne(ctx, filter, update, opts...)
	return err
}
