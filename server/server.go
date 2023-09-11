package server

import (
	"database/sql"
	"errors"
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strava-app/internal/api"
	"strava-app/internal/config"
	"strava-app/internal/db/repository"
	"strconv"
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
	// authentication middleware
	v1.Use(func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()
		userInfo, ok := header["X-User-Info"]
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON("missing user id in x-user-info request header")
		}
		id, _ := strconv.Atoi(userInfo)
		_, err := s.db.IsActiveUser(c.Context(), int64(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.Status(fiber.StatusUnauthorized).JSON("please sign up first")
			}
			msg := fmt.Sprintf("IsActiveUser failed: %s", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(msg)
		}
		c.Locals("user_id", id)
		return c.Next()
	})
	v1.Get("/activity/:id", s.route.GetActivity)

	// game
	v1.Post("/challenges", s.route.CreateChallenge)
	v1.Put("/challenges/:id", s.route.UpdateChallenge)
	v1.Get("/challenges/:id", s.route.GetChallenge)
	v1.Get("challenges", s.route.ListChallenge)
	v1.Get("/challenges/:id/gamers", s.route.ListGamerPerChallenge)
	v1.Get("/challenges/:id/longest-run-per-day", s.route.ListLongestRunPerActivity)
	v1.Put("/challenges/:id/join", s.route.JoinGame)
	v1.Delete("/challenges/:id/unjoin", s.route.UnJoinGame)

	return app.Listen(fmt.Sprintf(":%s", s.conf.ServicePort))
}

func (s *WebServer) ShutDown() error {
	return s.server.Shutdown()
}
