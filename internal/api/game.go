package api

import (
	"database/sql"
	"errors"
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strava-app/internal/db/repository/sqlite"
	"strava-app/internal/token"
	"strconv"
)

var adminGroup = map[int64]bool{
	121743168: true,
	113840436: true,
	112078641: true,
}

func (s handler) CreateChallenge(c *fiber.Ctx) error {
	payload := c.Locals(pasetoware.DefaultContextKey).(token.Claims)
	_, ok := adminGroup[payload.UserID]
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON("permission denied")
	}
	var params sqlite.CreateChallengeParams
	err := c.BodyParser(&params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	game, err := s.db.CreateChallenge(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

func (s handler) UpdateChallenge(c *fiber.Ctx) error {
	payload := c.Locals(pasetoware.DefaultContextKey).(token.Claims)
	_, ok := adminGroup[payload.UserID]
	if !ok {
		return c.Status(fiber.StatusForbidden).JSON("permission denied")
	}
	challengeID, _ := strconv.Atoi(c.Params("id"))
	var params sqlite.UpdateChallengeParams
	err := c.BodyParser(&params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	params.ID = int64(challengeID)
	game, err := s.db.UpdateChallenge(c.Context(), params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

func (s handler) GetChallenge(c *fiber.Ctx) error {
	challengeID, _ := strconv.Atoi(c.Params("id"))
	game, err := s.db.GetChallenge(c.Context(), int64(challengeID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

func (s handler) ListChallenge(c *fiber.Ctx) error {
	game, err := s.db.ListChallenge(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

func (s handler) ListGamerPerChallenge(c *fiber.Ctx) error {
	challengeID, _ := strconv.Atoi(c.Params("id"))
	game, err := s.db.ListGamers(c.Context(), int64(challengeID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

func (s handler) JoinGame(c *fiber.Ctx) error {
	var params sqlite.CreateGamerParams
	err := c.BodyParser(&params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	payload := c.Locals(pasetoware.DefaultContextKey).(token.Claims)
	challengeID, _ := strconv.Atoi(c.Params("id"))
	params.UserID = payload.UserID
	params.ChallengeID = int64(challengeID)
	game, err := s.db.CreateGamer(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

func (s handler) UnJoinGame(c *fiber.Ctx) error {
	payload := c.Locals(pasetoware.DefaultContextKey).(token.Claims)
	challengeID, _ := strconv.Atoi(c.Params("id"))
	err := s.db.DeleteGamer(c.Context(), sqlite.DeleteGamerParams{
		UserID:      payload.UserID,
		ChallengeID: int64(challengeID),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON("OK")
}

func (s handler) ListLongestRunPerActivity(c *fiber.Ctx) error {
	payload := c.Locals(pasetoware.DefaultContextKey).(token.Claims)
	game, err := s.db.GetLongestActivityPerDay(c.Context(), payload.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

type LongestRunPerDayQueryParams struct {
	Date string `json:"date"`
}

func (s handler) ListLongestRunPerActivity2(c *fiber.Ctx) error {

	var queryParams LongestRunPerDayQueryParams
	err := c.QueryParser(&queryParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	s.logger.Info("query params", zap.Reflect("params", queryParams))
	begin, end := GetDateTime(queryParams.Date)
	params := sqlite.GetLongestActivityPerDay2Params{
		Begin: begin,
		End:   end,
	}
	s.logger.Info("query with", zap.Int64("begin", begin), zap.Int64("end", end))
	game, err := s.db.GetLongestActivityPerDay2(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}
