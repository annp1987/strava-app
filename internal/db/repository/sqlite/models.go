// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package sqlite

import (
	"database/sql"
)

type Challenge struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Rules string `json:"rules"`
}

type Gamer struct {
	ChallengeID int64 `json:"challenge_id"`
	UserID      int64 `json:"user_id"`
	StartDate   int64 `json:"start_date"`
	EndDate     int64 `json:"end_date"`
	Target      int64 `json:"target"`
}

type RawActivity struct {
	ID             int64          `json:"id"`
	UserID         int64          `json:"user_id"`
	CreateAt       int64          `json:"create_at"`
	StartDate      int64          `json:"start_date"`
	StartDateLocal int64          `json:"start_date_local"`
	Distance       float64        `json:"distance"`
	AverageSpeed   float64        `json:"average_speed"`
	MovingTime     int64          `json:"moving_time"`
	Name           sql.NullString `json:"name"`
	SportType      string         `json:"sport_type"`
	MaxSpeed       float64        `json:"max_speed"`
	OriginalData   sql.NullString `json:"original_data"`
}

type RegisterUser struct {
	ID            int64          `json:"id"`
	UserName      string         `json:"user_name"`
	FirstName     sql.NullString `json:"first_name"`
	LastName      sql.NullString `json:"last_name"`
	ProfileMedium string         `json:"profile_medium"`
	Profile       string         `json:"profile"`
	AccessToken   string         `json:"access_token"`
	RefreshToken  string         `json:"refresh_token"`
	ExpiredAt     int64          `json:"expired_at"`
	Active        sql.NullInt64  `json:"active"`
}
