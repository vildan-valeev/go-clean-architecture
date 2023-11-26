package database_pg_test

import (
	"context"
	"os"
	"testing"

	"github.com/vildan-valeev/go-clean-architecture/pkg/database_pg"
)

func TestNewPooll(t *testing.T) {
	ctx := context.Background()
	dataSourceName := os.Getenv("TN_TEST_DSN")
	opts := []database_pg.Option{
		database_pg.WithLogLevel("debug"),
	}

	got, err := database_pg.NewPool(ctx, dataSourceName, opts...)
	if err != nil {
		t.Error(err)
		return
	}

	if got == nil {
		t.Error("NewPool() return nil")
	}
}
