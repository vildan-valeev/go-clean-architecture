package uuid_test

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/vildan-valeev/go-clean-architecture/pkg/database_pg/uuid"
)

func ExampleMain() {
	pgxConfig, err := pgxpool.ParseConfig(os.Getenv("PG_URI"))
	if err != nil {
		panic(err)
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())
		return nil
	}

	pgxConnPool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		panic(err)
	}

	// use pgxConnPool
	_ = pgxConnPool
}
