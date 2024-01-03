package database_mongo

import (
	"context"
	"github.com/rs/zerolog"
	"gitlab.com/tarataika/backend/pkg/logger/mongoadapter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Client struct {
	client *mongo.Client
	//level  string
}

func New() *Client {
	return &Client{client: nil}
}

func Connect(ctx context.Context, dsn, level string) (*mongo.Client, error) {
	serverAPI := options.
		ServerAPI(options.ServerAPIVersion1)
	zlogger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})
	loggerOpts := options.
		Logger().
		SetSink(mongoadapter.New(zlogger)).
		SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	opts := options.
		Client().
		ApplyURI(dsn).
		SetServerAPIOptions(serverAPI).
		SetServerSelectionTimeout(15 * time.Second).
		SetLoggerOptions(loggerOpts)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) Ping(ctx context.Context) error {
	return c.client.Ping(ctx, nil)
}

func (c *Client) Open(ctx context.Context, dsn, loglevel string) (err error) {
	c.client, err = Connect(ctx, dsn, loglevel)
	if err != nil {
		return err
	}

	return err
}

func (c *Client) Close(ctx context.Context) (err error) {
	if err = c.client.Disconnect(ctx); err != nil {
		//log.Debug("Debug message")
		log.Fatal()
	}
	return err
}

func (c *Client) Database(ctx context.Context, name string) *Database {
	db := c.client.Database(name)
	return &Database{
		mdb: db,
	}
}
