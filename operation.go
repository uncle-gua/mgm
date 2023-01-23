package mgm

import (
	"context"

	"github.com/uncle-uga/mgm/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func create(ctx context.Context, coll *Collection, model Model, opts ...*options.InsertOneOptions) error {
	// Call to saving hook
	if err := callToBeforeCreateHooks(ctx, model); err != nil {
		return err
	}

	res, err := coll.c.InsertOne(ctx, model, opts...)

	if err != nil {
		return err
	}

	// Set new id
	model.SetID(res.InsertedID)

	return callToAfterCreateHooks(ctx, model)
}

func first(ctx context.Context, coll *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return coll.c.FindOne(ctx, filter, opts...).Decode(model)
}

func update(ctx context.Context, coll *Collection, model Model, opts ...*options.UpdateOptions) error {
	// Call to saving hook
	if err := callToBeforeUpdateHooks(ctx, model); err != nil {
		return err
	}

	res, err := coll.c.UpdateOne(ctx, bson.M{field.ID: model.GetID()}, bson.M{"$set": model}, opts...)

	if err != nil {
		return err
	}

	return callToAfterUpdateHooks(ctx, res, model)
}

func del(ctx context.Context, coll *Collection, model Model) error {
	if err := callToBeforeDeleteHooks(ctx, model); err != nil {
		return err
	}
	res, err := coll.c.DeleteOne(ctx, bson.M{field.ID: model.GetID()})
	if err != nil {
		return err
	}

	return callToAfterDeleteHooks(ctx, res, model)
}
