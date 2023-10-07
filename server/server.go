package server

import (
	"encoding/json"
	"fmt"
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
	"strava-app/internal/api"
	"strava-app/internal/config"
	"strava-app/internal/db/repository"
	"strava-app/internal/token"
)

type WebServer struct {
	server *fiber.App
	conf   *config.Config
	db     repository.DBAPI
	route  api.ServerAPI
	logger *zap.Logger
}

func NewWebServer(config *config.Config, route api.ServerAPI, db repository.DBAPI, logger *zap.Logger) *WebServer {
	return &WebServer{
		server: fiber.New(),
		db:     db,
		conf:   config,
		route:  route,
		logger: logger,
	}
}

func (s *WebServer) Start() error {
	app := s.server
	app.Get("/connect", s.route.Connect)

	v1 := app.Group("/v1")
	v1.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE, PATCH",
		AllowHeaders:     "Accept, Content-Type, Authorization, Connection, Upgrade",
		AllowCredentials: true,
	}))
	// authentication middleware
	v1.Use(pasetoware.New(pasetoware.Config{
		SymmetricKey: []byte(token.SecretSymmetricKey),
		TokenPrefix:  "Bearer",
		Validate: func(decrypted []byte) (interface{}, error) {
			var payload token.Claims
			err := json.Unmarshal(decrypted, &payload)
			return payload, err
		},
	}))

	// activity
	v1.Get("/activity/:id", s.route.GetActivity)

	v1.Get("/me", s.route.GetMe)
	// user info
	v1.Get("/users/:id", s.route.GetUserInfo)
	// game
	v1.Post("/challenges", s.route.CreateChallenge)
	v1.Put("/challenges/:id", s.route.UpdateChallenge)
	v1.Get("/challenges/:id", s.route.GetChallenge)
	v1.Get("/challenges", s.route.ListChallenge)
	v1.Get("/challenges/:id/gamers", s.route.ListGamerPerChallenge)
	v1.Get("/challenges/:id/longest-run-per-day", s.route.ListLongestRunPerActivity)
	v1.Put("/challenges/:id/join", s.route.JoinGame)
	v1.Delete("/challenges/:id/unjoin", s.route.UnJoinGame)

	return app.Listen(fmt.Sprintf(":%s", s.conf.ServicePort))
}

func (s *WebServer) ShutDown() error {
	return s.server.Shutdown()
}
