package cronjob

import (
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"strava-app/internal/config"
	"strava-app/internal/db/repository"
	"time"
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
	s.c.AddFunc("@every 15s", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
		defer cancel()
		userIds, err := s.db.ListActiveUsers(ctx)
		if err != nil {
			s.logger.Error("ListActiveUsers failed: %s", zap.Error(err))
		}
		for _, id := range userIds {
			err := s.GetActivities(ctx, id)
			if err != nil {
				s.logger.Error("GetActivities for user failed: %s", zap.Error(err))
			}
		}
	})
	s.c.Run()
	return nil
}

func (s *CronServer) StopJob() error {
	s.c.Stop()
	return nil
}
