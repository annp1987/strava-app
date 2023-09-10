// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package sqlite

import (
	"context"
	_ "database/sql"
)

type Querier interface {
	CreateActivity(ctx context.Context, arg CreateActivityParams) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetToken(ctx context.Context, id int64) (GetTokenRow, error)
	UpdateToken(ctx context.Context, arg UpdateTokenParams) (int64, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (int64, error)
}

var _ Querier = (*Queries)(nil)