package api

import (
	"database/sql"
	"errors"
	pasetoware "github.com/gofiber/contrib/paseto"
	"github.com/gofiber/fiber/v2"
	"strava-app/internal/db/repository/sqlite"
	"strava-app/internal/token"
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

type UserInfo struct {
	ID            int64                           `json:"id"`
	UserName      string                          `json:"user_name"`
	FirstName     string                          `json:"first_name"`
	LastName      string                          `json:"last_name"`
	ProfileMedium string                          `json:"profile_medium"`
	Profile       string                          `json:"profile"`
	Challenges    []sqlite.GetJoinedChallengesRow `json:"challenges"`
}

func (s handler) GetMe(c *fiber.Ctx) error {
	payload := c.Locals(pasetoware.DefaultContextKey).(token.Claims)
	user, err := s.db.GetActiveUser(c.Context(), payload.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	challenges, err := s.db.GetJoinedChallenges(c.Context(), payload.UserID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
		}
	}
	userInfo := UserInfo{
		ID:            user.ID,
		UserName:      user.UserName,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		ProfileMedium: user.ProfileMedium,
		Profile:       user.Profile,
		Challenges:    challenges,
	}
	return c.JSON(userInfo)
}

func (s handler) GetUserInfo(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("id"))
	user, err := s.db.GetActiveUser(c.Context(), int64(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return c.JSON(user)
}
