package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/go-clean-architecture/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/vildan-valeev/go-clean-architecture/internal/config"
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
	categoryRepo "github.com/vildan-valeev/go-clean-architecture/internal/repository/category"
	itemRepo "github.com/vildan-valeev/go-clean-architecture/internal/repository/item"
	v1 "github.com/vildan-valeev/go-clean-architecture/internal/transport/http/v1"
	"github.com/vildan-valeev/go-clean-architecture/internal/transport/server"
	"github.com/vildan-valeev/go-clean-architecture/internal/usecase/category"
	"github.com/vildan-valeev/go-clean-architecture/internal/usecase/item"
	"github.com/vildan-valeev/go-clean-architecture/pkg/database_pg"
	redis "github.com/vildan-valeev/go-clean-architecture/pkg/database_redis"
)

func main() {
	logger.SetupLogging()

	log.Info().Msgf("Start App")

	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		log.Info().Msgf("Shutting down server. Reason: %s...", sig.String())
		cancel()
	}()

	// Instantiate a new type to represent our application.
	m := NewMain()

	// Execute program.
	if err := m.Run(ctx); err != nil {
		log.Error().Err(err).Msg("Run server error")

		_ = m.Close()

		os.Exit(1)
	}

	// Wait for CTRL-C.
	<-ctx.Done()

	// Clean up program.
	if err := m.Close(); err != nil {
		log.Error().Err(err).Msg("Shutting down server error")
		os.Exit(1)
	}

	log.Info().Msg("Bye!")
}

// Main represents the program.
type Main struct {
	// DB used by postgres service implementations.
	db db
	// Redis cache
	rs rs
	// HTTP server for handling communication.
	Srv *server.Server
}

// NewMain returns a new instance of Main.
func NewMain() *Main {
	return &Main{}
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (m *Main) Run(ctx context.Context) (err error) {
	cfg := config.NewConfig()
	logger.SetupLoggingLevel(cfg.LogLevel)

	m.db = database_pg.New(cfg.DSN, cfg.LogLevel)

	m.rs = redis.New(cfg.RedisHost, cfg.RedisPort)

	if err := m.db.Open(ctx); err != nil {
		return err
	}

	if err := m.rs.Open(ctx); err != nil {
		return err
	}

	if err := m.init(ctx, cfg); err != nil {
		return err
	}
	// Start the server.
	return m.Srv.Open()
}
func (m *Main) init(ctx context.Context, cfg *config.Config) error {
	it := itemRepo.New(m.db, m.rs)
	ct := categoryRepo.New(m.db)

	itemService := item.New(it)
	categoryService := category.New(ct)

	t := v1.NewTransport(v1.DI{
		Item:     itemService,
		Category: categoryService,
	})

	m.Srv = server.New(*cfg, t.Register())
	return nil
}

// Close gracefully stops the program.
func (m *Main) Close() (err error) {
	if m.Srv != nil {
		err = m.Srv.Close()
	}

	if m.db != nil {
		err = m.db.Close()
	}

	return err
}

type db interface {
	Open(context.Context) error
	Close() error

	repository.Database
}

type rs interface {
	Open(context.Context) error
	Close() error

	repository.RedisCache
}
