package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/redis/go-redis/v9"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain"
	mongoDatabase "github.com/vildan-valeev/go-clean-architecture/pkg/database_mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository - методы для работы с БД (интерфейс реализован в слое репо).
type CategoryRepo interface {
	PGInsertCategory(ctx context.Context, u domain.Category) error
	PGUpdateCategory(ctx context.Context, u domain.Category) error
	PGGetCategory(ctx context.Context, id uuid.UUID) (domain.Category, error)
	PGDeleteCategory(ctx context.Context, id uuid.UUID) error
}

// Repository - методы для работы с БД (интерфейс реализован в слое репо).
type ItemRepo interface {
	PGInsertItem(ctx context.Context, i domain.Item) error
	PGUpdateItem(ctx context.Context, i domain.Item) error
	PGGetItem(ctx context.Context, id uuid.UUID) (domain.Item, error)
	PGDeleteItem(ctx context.Context, id uuid.UUID) error

	RSInsertItem(ctx context.Context, i domain.Item) error
	RSUpdateItem(ctx context.Context, i domain.Item) error
	RSGetItem(ctx context.Context, id uuid.UUID) (domain.Item, error)
	RSDeleteItem(ctx context.Context, id uuid.UUID) error

	MGInsertItem(ctx context.Context, i domain.Item) error
	MGUpdateItem(ctx context.Context, i domain.Item) error
	MGGetItem(ctx context.Context, id uuid.UUID) (domain.Item, error)
	MGDeleteItem(ctx context.Context, id uuid.UUID) error
}

/*
Интерфейсы из инфры, удаленных сервисов, бд, и тд
*/
// ------------------------- POSTGRES --------------------------------------------------------
type (
	Tx             = pgx.Tx
	TxOptions      = pgx.TxOptions
	CommandTag     = pgconn.CommandTag
	Row            = pgx.Row
	Rows           = pgx.Rows
	Identifier     = pgx.Identifier
	CopyFromSource = pgx.CopyFromSource
	Batch          = pgx.Batch
	BatchResults   = pgx.BatchResults
)

type PG interface {
	Ping(context.Context) error

	Beginner
	TableOperator
}

// DB is the minimal representation of a relational database that powers repositories.
type PostgresDB interface {
	Beginner
	TableOperator
}

// TableOperator represents the methods required to interact with database
// tables. Typically this is satisfied by both databases and transactions.
type TableOperator interface {
	Execer
	Queryer
	Batcher
}

// Beginner represents an object that can begin a transaction.
type Beginner interface {
	Begin(context.Context) (Tx, error)
	BeginOpt(context.Context, TxOptions) (Tx, error)
}

// Execer executes a query that returns no result.
type Execer interface {
	Exec(context.Context, string, ...any) (CommandTag, error)
}

// Queryer executes a query that populates dest with the returned rows.
type Queryer interface {
	Query(context.Context, string, ...any) (Rows, error)
	QueryRow(context.Context, string, ...any) Row
}

type Batcher interface {
	CopyFrom(context.Context, Identifier, []string, CopyFromSource) (int64, error)
	SendBatch(context.Context, *Batch) BatchResults
}

// NotFound is a helper function to check if an error is `pgx.ErrNoRows`.
func NotFound(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

// ------------------------------------------ REDIS -----------------------------------------------

// RS In memory cache Redis.
type RedisDB interface {
	RedisQuery
	//RedisCommand
}
type RedisQuery interface {
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
}

// ----------------------------------- MONGO ----------------------------------------------------

type MongoDB interface {
	Collection(ctx context.Context, nameCollection string) *mongoDatabase.Collection
}

type MongoQuery interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}
