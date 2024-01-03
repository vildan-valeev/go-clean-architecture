package database_mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	mdb *mongo.Database
}

func (d *Database) Collection(ctx context.Context, nameCollection string) *Collection {
	c := d.mdb.Collection(nameCollection)
	return &Collection{
		coll: c,
	}
}
