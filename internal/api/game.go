package api

import (
	"github.com/gofiber/fiber/v2"
	"strava-app/internal/db/repository/sqlite"
)

func (s handler) JoinGame(c *fiber.Ctx) error {
	var params sqlite.CreateGamerParams
	err := c.BodyParser(&params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	id := c.Locals("user_id").(int)
	params.UserID = int64(id)
	_, err = s.db.CreateGamer(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(map[string]interface{}{"user_id": id})
}

func (s handler) UnJoinGame(c *fiber.Ctx) error {
	id := c.Locals("user_id").(int)
	err := s.db.DeleteGamer(c.Context(), int64(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON("OK")
}

func (s handler) ListGamerActivity(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
