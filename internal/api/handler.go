package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strava-app/internal/config"
	"strava-app/internal/db/repository"
)

type handler struct {
	conf   *config.Config
	db     repository.DBAPI
	logger *zap.Logger
}

type ServerAPI interface {
	Connect(c *fiber.Ctx) error
}

func NewServerAPI(config *config.Config, db repository.DBAPI, logger *zap.Logger) ServerAPI {
	return &handler{
		conf:   config,
		db:     db,
		logger: logger,
	}
}
