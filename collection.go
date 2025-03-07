package mgm

import (
	"context"

	"github.com/uncle-gua/mgm/builder"
	"github.com/uncle-gua/mgm/field"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection performs operations on models and the given Mongodb collection
type Collection struct {
	c *mongo.Collection
}

// FindByID method finds a doc and decodes it to a model, otherwise returns an error.
// The id field can be any value that if passed to the `PrepareID` method, it returns
// a valid ID (e.g string, bson.ObjectId).
func (coll *Collection) FindByID(id interface{}, model Model, opts ...*options.FindOneOptions) error {
	ctx, cancel := ctx()
	defer cancel()

	return coll.FindByIDWithCtx(ctx, id, model, opts...)
}

// FindByIDWithCtx method finds a doc and decodes it to a model, otherwise returns an error.
// The id field can be any value that if passed to the `PrepareID` method, it returns
// a valid ID (e.g string, bson.ObjectId).
func (coll *Collection) FindByIDWithCtx(ctx context.Context, id interface{}, model Model, opts ...*options.FindOneOptions) error {
	id, err := model.PrepareID(id)

	if err != nil {
		return err
	}

	return first(ctx, coll, bson.M{field.ID: id}, model, opts...)
}

// First method searches and returns the first document in the search results.
func (coll *Collection) First(filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	ctx, cancel := ctx()
	defer cancel()

	return coll.FirstWithCtx(ctx, filter, model, opts...)
}

// FirstWithCtx method searches and returns the first document in the search results.
func (coll *Collection) FirstWithCtx(ctx context.Context, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return first(ctx, coll, filter, model, opts...)
}

// Create method inserts a new model into the database.
func (coll *Collection) Create(model Model, opts ...*options.InsertOneOptions) error {
	ctx, cancel := ctx()
	defer cancel()

	return coll.CreateWithCtx(ctx, model, opts...)
}

// CreateWithCtx method inserts a new model into the database.
func (coll *Collection) CreateWithCtx(ctx context.Context, model Model, opts ...*options.InsertOneOptions) error {
	return create(ctx, coll, model, opts...)
}

// Update function persists the changes made to a model to the database.
// Calling this method also invokes the model's mgm updating, updated,
// saving, and saved hooks.
func (coll *Collection) Update(model Model, opts ...*options.UpdateOptions) error {
	ctx, cancel := ctx()
	defer cancel()

	return coll.UpdateWithCtx(ctx, model, opts...)
}

// UpdateWithCtx function persists the changes made to a model to the database using the specified context.
// Calling this method also invokes the model's mgm updating, updated,
// saving, and saved hooks.
func (coll *Collection) UpdateWithCtx(ctx context.Context, model Model, opts ...*options.UpdateOptions) error {
	return update(ctx, coll, model, opts...)
}

// Update function persists the changes made to a model to the database.
// Calling this method also invokes the model's mgm updating, updated,
// saving, and saved hooks.
func (coll *Collection) UpdateByID(id interface{}, model Model, opts ...*options.UpdateOptions) error {
	ctx, cancel := ctx()
	defer cancel()

	return coll.UpdateByIDWithCtx(ctx, id, model, opts...)
}

// UpdateWithCtx function persists the changes made to a model to the database using the specified context.
// Calling this method also invokes the model's mgm updating, updated,
// saving, and saved hooks.
func (coll *Collection) UpdateByIDWithCtx(ctx context.Context, id interface{}, model Model, opts ...*options.UpdateOptions) error {
	id, err := model.PrepareID(id)

	if err != nil {
		return err
	}
	model.SetID(id)

	return update(ctx, coll, model, opts...)
}

// Delete method deletes a model (doc) from a collection.
// To perform additional operations when deleting a model
// you should use hooks rather than overriding this method.
func (coll *Collection) Delete(model Model) error {
	ctx, cancel := ctx()
	defer cancel()

	return del(ctx, coll, model)
}

// Name method return the collection name.
func (coll *Collection) Name() string {
	return coll.c.Name()
}

func (coll *Collection) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	ctx, cancel := ctx()
	defer cancel()

	return coll.CountDocumentsWithCtx(ctx, filter, opts...)
}

func (coll *Collection) CountDocumentsWithCtx(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return coll.c.CountDocuments(ctx, filter, opts...)
}

func (coll *Collection) Find(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	ctx, cancel := ctx()
	defer cancel()

	return coll.FindWithCtx(ctx, filter, opts...)
}

func (coll *Collection) FindWithCtx(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return coll.c.Find(ctx, filter)
}

func (coll *Collection) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := ctx()
	defer cancel()

	return coll.DeleteManyCtx(ctx, filter, opts...)
}

func (coll *Collection) DeleteManyCtx(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return coll.c.DeleteMany(ctx, filter)
}

func (coll *Collection) InsertMany(documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	ctx, cancel := ctx()
	defer cancel()

	return coll.InsertManyCtx(ctx, documents, opts...)
}

func (coll *Collection) InsertManyCtx(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return coll.c.InsertMany(ctx, documents, opts...)
}

func (coll *Collection) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, cancel := ctx()
	defer cancel()

	return coll.FindOneWithCtx(ctx, filter, opts...)
}

func (coll *Collection) FindOneWithCtx(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return coll.c.FindOne(ctx, filter, opts...)
}

func (coll *Collection) Aggregate(pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	ctx, cancel := ctx()
	defer cancel()

	return coll.AggregateWithCtx(ctx, pipeline, opts...)
}

func (coll *Collection) AggregateWithCtx(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return coll.c.Aggregate(ctx, pipeline, opts...)
}

// DeleteWithCtx method deletes a model (doc) from a collection using the specified context.
// To perform additional operations when deleting a model
// you should use hooks rather than overriding this method.
func (coll *Collection) DeleteWithCtx(ctx context.Context, model Model) error {
	return del(ctx, coll, model)
}

// SimpleFind finds, decodes and returns the results.
func (coll *Collection) SimpleFind(results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	ctx, cancel := ctx()
	defer cancel()

	return coll.SimpleFindWithCtx(ctx, results, filter, opts...)
}

// SimpleFindWithCtx finds, decodes and returns the results using the specified context.
func (coll *Collection) SimpleFindWithCtx(ctx context.Context, results interface{}, filter interface{}, opts ...*options.FindOptions) error {
	cur, err := coll.c.Find(ctx, filter, opts...)

	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

//--------------------------------
// Aggregation methods
//--------------------------------

// SimpleAggregateFirst is just same as SimpleAggregateFirstWithCtx, but doesn't get context param.
func (coll *Collection) SimpleAggregateFirst(result interface{}, stages ...interface{}) (bool, error) {
	ctx, cancel := ctx()
	defer cancel()

	return coll.SimpleAggregateFirstWithCtx(ctx, result, stages...)
}

// SimpleAggregateFirstWithCtx performs a simple aggregation, decodes the first aggregate result and returns it using the provided result parameter.
// The value of `stages` can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (coll *Collection) SimpleAggregateFirstWithCtx(ctx context.Context, result interface{}, stages ...interface{}) (bool, error) {
	cur, err := coll.SimpleAggregateCursorWithCtx(ctx, stages...)
	if err != nil {
		return false, err
	}
	if cur.Next(ctx) {
		return true, cur.Decode(result)
	}
	return false, nil
}

// SimpleAggregate is just same as SimpleAggregateWithCtx, but doesn't get context param.
func (coll *Collection) SimpleAggregate(results interface{}, stages ...interface{}) error {
	ctx, cancel := ctx()
	defer cancel()

	return coll.SimpleAggregateWithCtx(ctx, results, stages...)
}

// SimpleAggregateWithCtx performs a simple aggregation, decodes the aggregate result and returns the list using the provided result parameter.
// The value of `stages` can be Operator|bson.M
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (coll *Collection) SimpleAggregateWithCtx(ctx context.Context, results interface{}, stages ...interface{}) error {
	cur, err := coll.SimpleAggregateCursorWithCtx(ctx, stages...)
	if err != nil {
		return err
	}

	return cur.All(ctx, results)
}

// SimpleAggregateCursor is just same as SimpleAggregateCursorWithCtx, but
// doesn't get context.
func (coll *Collection) SimpleAggregateCursor(stages ...interface{}) (*mongo.Cursor, error) {
	ctx, cancel := ctx()
	defer cancel()

	return coll.SimpleAggregateCursorWithCtx(ctx, stages...)
}

// SimpleAggregateCursorWithCtx performs a simple aggregation and returns a cursor over the resulting documents.
// Note: you can not use this method in a transaction because it does not accept a context.
// To participate in transactions, please use the regular aggregation method.
func (coll *Collection) SimpleAggregateCursorWithCtx(ctx context.Context, stages ...interface{}) (*mongo.Cursor, error) {
	pipeline := bson.A{}

	for _, stage := range stages {
		if operator, ok := stage.(builder.Operator); ok {
			pipeline = append(pipeline, builder.S(operator))
		} else {
			pipeline = append(pipeline, stage)
		}
	}

	return coll.c.Aggregate(ctx, pipeline, nil)
}
