// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package sqlite

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT OR REPLACE INTO register_users (
    id, user_name, first_name, last_name, profile_medium, profile, access_token, refresh_token, expired_at
) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?)
    RETURNING id
`

type CreateUserParams struct {
	ID            int64  `json:"id"`
	UserName      string `json:"user_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ProfileMedium string `json:"profile_medium"`
	Profile       string `json:"profile"`
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	ExpiredAt     int64  `json:"expired_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.UserName,
		arg.FirstName,
		arg.LastName,
		arg.ProfileMedium,
		arg.Profile,
		arg.AccessToken,
		arg.RefreshToken,
		arg.ExpiredAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const disableUser = `-- name: DisableUser :one
UPDATE register_users
set active = ?
WHERE id = ?
    RETURNING id
`

type DisableUserParams struct {
	Active sql.NullInt64 `json:"active"`
	ID     int64         `json:"id"`
}

func (q *Queries) DisableUser(ctx context.Context, arg DisableUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, disableUser, arg.Active, arg.ID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getActiveUser = `-- name: GetActiveUser :one
SELECT id, user_name, first_name, last_name, profile_medium, profile
FROM register_users
WHERE id = ? AND active = 1
`

type GetActiveUserRow struct {
	ID            int64  `json:"id"`
	UserName      string `json:"user_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ProfileMedium string `json:"profile_medium"`
	Profile       string `json:"profile"`
}

func (q *Queries) GetActiveUser(ctx context.Context, id int64) (GetActiveUserRow, error) {
	row := q.db.QueryRowContext(ctx, getActiveUser, id)
	var i GetActiveUserRow
	err := row.Scan(
		&i.ID,
		&i.UserName,
		&i.FirstName,
		&i.LastName,
		&i.ProfileMedium,
		&i.Profile,
	)
	return i, err
}

const getJoinedChallenges = `-- name: GetJoinedChallenges :many
SELECT c.id, c.name
FROM gamers AS g
 JOIN challenges AS c ON g.challenge_id = c.id
WHERE g.user_id = ?
`

type GetJoinedChallengesRow struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) GetJoinedChallenges(ctx context.Context, userID int64) ([]GetJoinedChallengesRow, error) {
	rows, err := q.db.QueryContext(ctx, getJoinedChallenges, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetJoinedChallengesRow{}
	for rows.Next() {
		var i GetJoinedChallengesRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const getToken = `-- name: GetToken :one
SELECT access_token, refresh_token, expired_at FROM register_users
WHERE id = ? LIMIT 1
`

type GetTokenRow struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
}

func (q *Queries) GetToken(ctx context.Context, id int64) (GetTokenRow, error) {
	row := q.db.QueryRowContext(ctx, getToken, id)
	var i GetTokenRow
	err := row.Scan(&i.AccessToken, &i.RefreshToken, &i.ExpiredAt)
	return i, err
}

const isActiveUser = `-- name: IsActiveUser :one
SELECT id
FROM register_users
WHERE id = ? AND active = 1
`

func (q *Queries) IsActiveUser(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, isActiveUser, id)
	err := row.Scan(&id)
	return id, err
}

const listActiveUsers = `-- name: ListActiveUsers :many
SELECT id
FROM register_users
WHERE active = 1
`

func (q *Queries) ListActiveUsers(ctx context.Context) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, listActiveUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateToken = `-- name: UpdateToken :one
UPDATE register_users
set access_token = ?, refresh_token = ?, expired_at = ?
WHERE id = ?
    RETURNING id
`

type UpdateTokenParams struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiredAt    int64  `json:"expired_at"`
	ID           int64  `json:"id"`
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateToken,
		arg.AccessToken,
		arg.RefreshToken,
		arg.ExpiredAt,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}
