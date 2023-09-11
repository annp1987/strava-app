package api

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strava-app/internal/db/repository/sqlite"
	"strconv"
)

func (s handler) CreateChallenge(c *fiber.Ctx) error {
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
	id := c.Locals("user_id").(int)
	challengeID, _ := strconv.Atoi(c.Params("id"))
	params.UserID = int64(id)
	params.ChallengeID = int64(challengeID)
	game, err := s.db.CreateGamer(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}

func (s handler) UnJoinGame(c *fiber.Ctx) error {
	id := c.Locals("user_id").(int)
	challengeID, _ := strconv.Atoi(c.Params("id"))
	err := s.db.DeleteGamer(c.Context(), sqlite.DeleteGamerParams{
		UserID:      int64(id),
		ChallengeID: int64(challengeID),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON("OK")
}

func (s handler) ListLongestRunPerActivity(c *fiber.Ctx) error {
	challengeID, _ := strconv.Atoi(c.Params("id"))
	game, err := s.db.GetLongestActivityPerDay(c.Context(), int64(challengeID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(game)
}
