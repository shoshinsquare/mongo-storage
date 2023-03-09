package nftmodel

import (
	"context"
	"time"

	"github.com/ponlv/go-kit/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(ctx context.Context, data *NFT) (interface{}, error) {
	collection := mongodb.Coll(data)

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	id, err := collection.CreateWithCtx(ctx, data)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func FindWithObjectID(ctx context.Context, objectID string) (*NFT, error) {
	filter := bson.M{"object_id": objectID}
	return findWithCondition(ctx, filter)
}

func FindById(ctx context.Context, id string) (*NFT, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	return findWithCondition(ctx, filter)
}

func findWithCondition(ctx context.Context, filter interface{}, findOptions ...*options.FindOneOptions) (*NFT, error) {
	coll := mongodb.CollRead(&NFT{})

	result := &NFT{}
	if err := coll.FirstWithCtx(ctx, filter, result, findOptions...); err != nil {
		return nil, err
	}

	return result, nil
}
