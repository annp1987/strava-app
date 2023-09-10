package server

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strava-app/internal/api"
	"strava-app/internal/config"
)

type WebServer struct {
	Server *fiber.App
	Conf   *config.Config
	route  api.ServerAPI
	logger *zap.Logger
}

func NewWebServer(config *config.Config, route api.ServerAPI, logger *zap.Logger) *WebServer {
	return &WebServer{
		Server: fiber.New(),
		Conf:   config,
		route:  route,
		logger: logger,
	}
}

func (s *WebServer) Start() error {
	app := s.Server
	app.Get("/connect", s.route.Connect)

	return app.Listen(fmt.Sprintf(":%s", s.Conf.ServicePort))
}

func (s *WebServer) ShutDown() error {
	return s.Server.Shutdown()
}
