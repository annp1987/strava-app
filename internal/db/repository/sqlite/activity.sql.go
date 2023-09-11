// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: activity.sql

package sqlite

import (
	"context"
	"database/sql"
)

const createActivity = `-- name: CreateActivity :exec
INSERT OR IGNORE INTO raw_activities (
    id, user_id, create_at, start_date, start_date_local, distance, average_speed, moving_time, name, sport_type, max_speed, original_data
) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateActivityParams struct {
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

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) error {
	_, err := q.db.ExecContext(ctx, createActivity,
		arg.ID,
		arg.UserID,
		arg.CreateAt,
		arg.StartDate,
		arg.StartDateLocal,
		arg.Distance,
		arg.AverageSpeed,
		arg.MovingTime,
		arg.Name,
		arg.SportType,
		arg.MaxSpeed,
		arg.OriginalData,
	)
	return err
}

const getActivity = `-- name: GetActivity :many
SELECT
    id,
    user_id,
    create_at,
    start_date,
    start_date_local,
    distance,
    average_speed,
    moving_time,
    name,
    sport_type,
    max_speed
FROM raw_activities
WHERE user_id=? AND sport_type=?
`

type GetActivityParams struct {
	UserID    int64  `json:"user_id"`
	SportType string `json:"sport_type"`
}

type GetActivityRow struct {
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
}

func (q *Queries) GetActivity(ctx context.Context, arg GetActivityParams) ([]GetActivityRow, error) {
	rows, err := q.db.QueryContext(ctx, getActivity, arg.UserID, arg.SportType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetActivityRow{}
	for rows.Next() {
		var i GetActivityRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CreateAt,
			&i.StartDate,
			&i.StartDateLocal,
			&i.Distance,
			&i.AverageSpeed,
			&i.MovingTime,
			&i.Name,
			&i.SportType,
			&i.MaxSpeed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
