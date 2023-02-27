package usermodel

import (
	"context"
	"github.com/ponlv/go-kit/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(ctx context.Context, data *User) (interface{}, error) {
	collection := mongodb.Coll(data)
	id, err := collection.CreateWithCtx(ctx, data)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func FindWithEmail(ctx context.Context, email string) (*User, error) {
	filter := bson.M{"email": email}
	return findWithCondition(ctx, filter)
}

func FindById(ctx context.Context, id string) (*User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	return findWithCondition(ctx, filter)
}

func findWithCondition(ctx context.Context, filter interface{}, findOptions ...*options.FindOneOptions) (*User, error) {
	coll := mongodb.CollRead(&User{})

	result := &User{}
	if err := coll.FirstWithCtx(ctx, filter, result, findOptions...); err != nil {
		return nil, err
	}

	return result, nil
}
