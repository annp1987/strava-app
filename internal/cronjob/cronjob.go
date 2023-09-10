package cronjob

import (
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"strava-app/internal/config"
	"strava-app/internal/db/repository"
	"time"
)

const (
	ClubID = 1116587
	UserID = 112078641
)

type CronServer struct {
	conf   *config.Config
	c      *cron.Cron
	db     repository.DBAPI
	logger *zap.Logger
}

type JobInterface interface {
	StartJob() error
	StopJob() error
}

func NewCronJob(conf *config.Config, db repository.DBAPI, logger *zap.Logger) *CronServer {
	return &CronServer{
		conf:   conf,
		c:      cron.New(),
		db:     db,
		logger: logger,
	}
}

func (s *CronServer) StartJob() error {
	s.c.AddFunc("@every 30m", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		s.GetActivities(ctx, 112078641)
	})
	s.c.Run()
	return nil
}

func (s *CronServer) StopJob() error {
	s.c.Stop()
	return nil
}
