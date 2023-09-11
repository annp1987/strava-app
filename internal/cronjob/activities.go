package cronjob

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	strava "github.com/obalunenko/strava-api/client"
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
	var after = time.Now().Add(-15 * 24 * time.Hour).Unix()
	activities, err := apiClient.Activities.GetLoggedInAthleteActivities(ctx, strava.GetLoggedInAthleteActivitiesOpts{
		After: &after,
	})
	if err != nil {
		return fmt.Errorf("GetLoggedInAthleteActivities failed: %s", err.Error())
	}

	if len(activities) > 0 {
		var params = make([]sqlite.CreateActivityParams, 0)
		for _, act := range activities {
			currentTime := time.Now().Unix()
			startDate := time.Time(act.StartDate).Unix()
			param := sqlite.CreateActivityParams{
				ID:             act.ID,
				UserID:         act.Athlete.ID,
				CreateAt:       currentTime,
				StartDate:      startDate,
				StartDateLocal: time.Time(act.StartDateLocal).Unix(),
				Distance:       float64(act.Distance),
				AverageSpeed:   float64(act.AverageSpeed),
				MovingTime:     act.MovingTime,
				Name:           sql.NullString{String: act.Name, Valid: true},
				SportType:      string(act.SportType),
				MaxSpeed:       float64(act.MaxSpeed),
				OriginalData:   sql.NullString{},
			}
			params = append(params, param)
		}
		if err = s.db.CreateActivityTx(ctx, params); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil
			}
			return fmt.Errorf("CreateActivityTx failed: %s", err)
		}
	}
	return nil
}
