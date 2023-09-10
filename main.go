package main

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"strava-app/internal/api"
	"strava-app/internal/config"
	"strava-app/internal/cronjob"
	"strava-app/internal/db"
	"strava-app/internal/db/repository"
	"strava-app/internal/loggerfx"
	"strava-app/server"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			loggerfx.ProvideLogger,
			db.NewDBClient,
			server.NewWebServer,
			cronjob.NewCronJob,
			repository.NewRepo,
			api.NewServerAPI),
		fx.Invoke(RegisterWebServer),
		fx.Invoke(RegisterCronJobServer))
	app.Run()
}

func RegisterWebServer(lifeCycle fx.Lifecycle, webServer *server.WebServer, conf *config.Config, logger *zap.Logger) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := webServer.Start(); err != nil {
					logger.Fatal("start server error : ", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Info("stopping server ...")
			return webServer.ShutDown()
		},
	})
	logger.Info("web server started on ", zap.String("port", conf.ServicePort))
}

func RegisterCronJobServer(lifeCycle fx.Lifecycle, cronjob *cronjob.CronServer, conf *config.Config, logger *zap.Logger) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := cronjob.StartJob(); err != nil {
					logger.Fatal("start server error : ", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Info("stopping server ...")
			return cronjob.StopJob()
		},
	})
	logger.Info("web server started on ", zap.String("port", conf.ServicePort))
}
