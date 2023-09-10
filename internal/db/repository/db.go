package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strava-app/internal/config"
	"strava-app/internal/db/repository/sqlite"
	"time"
)

const AuthURL = "https://www.strava.com/api/v3/oauth/token"

type Token struct {
	TokenType    string `json:"token_type"`
	ExpiresAt    int64  `json:"expires_at"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type DBAPI interface {
	sqlite.Querier
	GetUserToken(ctx context.Context, id int64) (string, error)
	CreateActivityTx(ctx context.Context, activity []sqlite.CreateActivityParams) error
}

type database struct {
	conf *config.Config
	db   *sql.DB
	*sqlite.Queries
	logger *zap.Logger
}

func NewRepo(db *sql.DB, conf *config.Config, logger *zap.Logger) DBAPI {
	return &database{conf: conf, db: db, Queries: sqlite.New(db), logger: logger}
}

func (s *database) GetUserToken(ctx context.Context, id int64) (string, error) {
	token, err := s.GetToken(ctx, id)
	if err != nil {
		return "", fmt.Errorf("GetToken error: %s", err.Error())
	}
	// check token is expired or not
	if token.ExpiredAt <= time.Now().Unix() {
		// request new token
		oauthURI := fmt.Sprintf(
			"%s?client_id=%s&client_secret=%s&grant_type=%s&refresh_token=%s",
			AuthURL,
			s.conf.ClientID,
			s.conf.ClientSecret,
			"refresh_token",
			token.RefreshToken,
		)
		resp, err := http.Post(oauthURI, "application/json", nil)
		if err != nil {
			return "", fmt.Errorf("post request failed: %s", err.Error())
		}
		defer resp.Body.Close()
		var refreshToken Token
		jsonData, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(jsonData, &refreshToken)
		if err != nil {
			return "", fmt.Errorf("read response failed: %s", err.Error())
		}
		//update token
		params := sqlite.UpdateTokenParams{
			AccessToken:  refreshToken.AccessToken,
			RefreshToken: refreshToken.RefreshToken,
			ExpiredAt:    refreshToken.ExpiresAt,
			ID:           id,
		}
		_, err = s.UpdateToken(ctx, params)
		if err != nil {
			return "", fmt.Errorf("UpdateToken failed: %s", err.Error())
		}
		return refreshToken.AccessToken, nil
	}
	return token.AccessToken, nil
}

func (s *database) CreateActivityTx(ctx context.Context, activity []sqlite.CreateActivityParams) error {
	err := s.execTx(ctx, func(queries *sqlite.Queries) error {
		for _, act := range activity {
			_, err := queries.CreateActivity(ctx, act)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (s *database) execTx(ctx context.Context, fn func(*sqlite.Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := sqlite.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
