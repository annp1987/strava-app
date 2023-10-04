package api

import (
	"github.com/gofiber/fiber/v2"
	"strava-app/internal/db/repository/sqlite"
	"strconv"
)

func (s handler) GetActivity(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("id"))
	params := sqlite.GetActivityParams{
		UserID:    int64(userID),
		SportType: "Run",
	}
	activities, err := s.db.GetActivity(c.Context(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(activities)
}

func (s handler) GetUserInfo(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("id"))
	user, err := s.db.GetActiveUser(c.Context(), int64(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(user)
}
