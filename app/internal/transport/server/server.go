// Package httpserver implements HTTP server.
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

	_ "github.com/swaggo/fiber-swagger/example/docs"
	"github.com/vildan-valeev/go-clean-architecture/internal/config"
	"net"
)

type Server struct {
	http   *fiber.App
	config config.Config
}

// New returns a new instance of Server.
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
	//docs.SwaggerInfo.Title = "Clean Architecture API"
	//docs.SwaggerInfo.Description = "This is a sample server."
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "localhost:8000"
	//docs.SwaggerInfo.BasePath = ""
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}

	s.http.Use(favicon.New())
	s.http.Use(requestid.New())
	s.http.Use(logger.Middleware())
	s.http.Use(cors.New())
	s.http.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	//s.http.Get("/swagger/*", swagger.HandlerDefault) // default

	s.http.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:8000/swagger/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		//OAuth: &swagger.OAuthConfig{
		//	AppName:  "OAuth Provider",
		//	ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		//},
		//// Ability to change OAuth2 redirect uri location
		//OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	// TODO: s.http.Group...
	s.http.Mount("/", handlers)

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
