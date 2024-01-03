package database_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	coll *mongo.Collection
}

func (db *Collection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return db.coll.InsertOne(ctx, document, opts...)
}

func (db *Collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return db.coll.FindOne(ctx, filter, opts...)
}

func (db *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return db.coll.UpdateOne(ctx, filter, update, opts...)
}

func (db *Collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	return db.coll.Find(ctx, filter, opts...)
}
