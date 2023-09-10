package cronjob

import (
	"context"
	"database/sql"
	strava "github.com/obalunenko/strava-api/client"
	"go.uber.org/zap"
	"strava-app/internal/db/repository/sqlite"
	"time"
)

func (s *CronServer) GetActivities(ctx context.Context, id int64) error {
	token, err := s.db.GetUserToken(ctx, id)
	if err != nil {
		return err
	}
	apiClient, err := strava.NewAPIClient(token)
	if err != nil {
		return err
	}
	var after = time.Now().Add(-30 * time.Minute).Unix()
	activities, err := apiClient.Activities.GetLoggedInAthleteActivities(ctx, strava.GetLoggedInAthleteActivitiesOpts{
		After: &after,
	})
	if len(activities) > 0 {
		s.logger.Info("activity", zap.Reflect("ctx", activities[0]))
	}

	var params = make([]sqlite.CreateActivityParams, 0)
	for _, act := range activities {
		param := sqlite.CreateActivityParams{
			ID:           act.ID,
			UserID:       act.Athlete.ID,
			CreateAt:     time.Now().Unix(),
			StartDate:    time.Time(act.StartDate).Unix(),
			Distance:     float64(act.Distance),
			AverageSpeed: float64(act.AverageSpeed),
			MovingTime:   act.MovingTime,
			Name:         sql.NullString{String: act.Name, Valid: true},
			SportType:    string(act.SportType),
			MaxSpeed:     float64(act.MaxSpeed),
			OriginalData: sql.NullString{},
		}
		params = append(params, param)
	}
	return s.db.CreateActivityTx(ctx, params)
}
