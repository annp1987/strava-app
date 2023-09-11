package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strava-app/internal/db/repository/sqlite"
	"strava-app/internal/token"
	"time"
)

const AuthURL = "https://www.strava.com/api/v3/oauth/token"

type Athlete struct {
	ID            int       `json:"id"`
	Username      string    `json:"username"`
	ResourceState int       `json:"resource_state"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	Sex           string    `json:"sex"`
	Premium       bool      `json:"premium"`
	Summit        bool      `json:"summit"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	BadgeTypeID   int       `json:"badge_type_id"`
	ProfileMedium string    `json:"profile_medium"`
	Profile       string    `json:"profile"`
	Friend        int       `json:"friend"`
	Follower      int       `json:"follower"`
}

type OauthResponse struct {
	TokenType    string  `json:"token_type"`
	ExpiresAt    int64   `json:"expires_at"`
	ExpiresIn    int64   `json:"expires_in"`
	RefreshToken string  `json:"refresh_token"`
	AccessToken  string  `json:"access_token"`
	Athlete      Athlete `json:"athlete"`
}

func (s handler) Connect(c *fiber.Ctx) error {
	var oathResp OauthResponse

	// build token url
	oauthURI := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s&grant_type=%s",
		AuthURL,
		s.conf.ClientID,
		s.conf.ClientSecret,
		c.Query("code"),
		"authorization_code",
	)
	resp, err := http.Post(oauthURI, "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	jsonData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonData, &oathResp)
	if err != nil {
		fmt.Println(err)
	}
	s.logger.Info("athlete", zap.Reflect("user", oathResp.Athlete))
	params := sqlite.CreateUserParams{
		ID:            int64(oathResp.Athlete.ID),
		UserName:      oathResp.Athlete.Username,
		FirstName:     oathResp.Athlete.FirstName,
		LastName:      oathResp.Athlete.LastName,
		ProfileMedium: oathResp.Athlete.ProfileMedium,
		Profile:       oathResp.Athlete.Profile,
		AccessToken:   oathResp.AccessToken,
		RefreshToken:  oathResp.RefreshToken,
		ExpiredAt:     oathResp.ExpiresAt,
	}

	// Generate encoded token and send it as response.
	t, err := token.GenerateToken(int64(oathResp.Athlete.ID))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println("token: ", t)
	if _, err = s.db.CreateUser(context.Background(), params); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "create user %s", err.Error())
	}
	redirectUrl := fmt.Sprintf("%s?id=%d&user_name=%s&profile=%s&profile_medium=%s&token=%s", s.conf.RedirectURL,
		oathResp.Athlete.ID, oathResp.Athlete.Username, oathResp.Athlete.Profile, oathResp.Athlete.ProfileMedium, t)
	return c.Redirect(redirectUrl, 301)
}
