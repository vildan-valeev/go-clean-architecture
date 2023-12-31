package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"
	"github.com/vildan-valeev/go-clean-architecture/pkg/logger"

	"github.com/vildan-valeev/go-clean-architecture/internal/config"
	"net"
)

type Server struct {
	http   *fiber.App
	config config.Config
}

func New(cfg config.Config, handlers *fiber.App) *Server {
	s := &Server{
		config: cfg,
	}

	s.http = fiber.New(fiber.Config{
		ServerHeader:          "Order Service",
		DisableStartupMessage: true,
		DisableKeepalive:      true,
	})
	// TODO: s.http.Group...

	s.http.Use(favicon.New())
	s.http.Use(requestid.New())
	s.http.Use(logger.Middleware())
	s.http.Use(cors.New())
	s.http.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	s.http.Get("/swagger/*", swagger.HandlerDefault) // default

	// TODO: s.http.Group...
	s.http.Mount("/v1", handlers)

	return s
}

// Open validates the server options and begins listening on the bind address.
func (s *Server) Open() error {

	go func() {
		address := net.JoinHostPort(s.config.IP, s.config.HTTPPort)
		log.Info().Msgf("Start HTTP on %s", address)

		if err := s.http.Listen(address); err != nil {
			log.Fatal().Err(err).Msg("failed to http serve")
		}
	}()

	return nil
}

// Close gracefully shuts down the server.
func (s *Server) Close() error {
	return s.http.Shutdown()
}
