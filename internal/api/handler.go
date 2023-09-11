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
	GetActivity(c *fiber.Ctx) error
	CreateChallenge(c *fiber.Ctx) error
	UpdateChallenge(c *fiber.Ctx) error
	GetChallenge(c *fiber.Ctx) error
	ListChallenge(c *fiber.Ctx) error
	ListGamerPerChallenge(c *fiber.Ctx) error
	ListLongestRunPerActivity(c *fiber.Ctx) error
	JoinGame(c *fiber.Ctx) error
	UnJoinGame(c *fiber.Ctx) error
}

func NewServerAPI(config *config.Config, db repository.DBAPI, logger *zap.Logger) ServerAPI {
	return &handler{
		conf:   config,
		db:     db,
		logger: logger,
	}
}
