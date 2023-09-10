package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"strava-app/internal/config"
)

func NewDBClient(conf *config.Config, logger *zap.Logger) *sql.DB {
	fmt.Println("NewSqlDBClient")
	dbtx, err := sql.Open("sqlite3", conf.SQLiteDBFile)
	if err != nil {
		logger.Error("Database connection error", zap.Error(err))
		panic(err)
	}
	logger.Info("Successful connected to sqlite3 db")
	return dbtx
}
