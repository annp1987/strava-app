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
		today := time.Now().Truncate(24 * time.Hour).Unix()
		lr, err := s.db.GetCurrentLongestRunPerDay(ctx, sqlite.GetCurrentLongestRunPerDayParams{
			UserID: id,
			Today:  today,
		})
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("GetCurrentLongestRunPerDay failed: %s", err.Error())
			}
			lr = sqlite.LongestRunPerDay{
				UserID:       id,
				Today:        today,
				ActivityID:   0,
				StartDate:    0,
				Distance:     0,
				AverageSpeed: 0,
				MovingTime:   0,
				Name:         sql.NullString{},
				SportType:    "",
				MaxSpeed:     0,
			}
		}
		var params = make([]sqlite.CreateActivityParams, 0)
		for _, act := range activities {
			currentTime := time.Now().Unix()
			startDate := time.Time(act.StartDate).Unix()
			// may be valid longest run
			if (startDate+act.MovingTime)+3600 <= currentTime {
				if lr.Distance <= float64(act.Distance) && act.SportType == "Run" {
					lr.ActivityID = act.Athlete.ID
					lr.MovingTime = act.MovingTime
					lr.StartDate = startDate
					lr.Name = sql.NullString{String: act.Name, Valid: true}
					lr.Distance = float64(act.Distance)
					lr.AverageSpeed = float64(act.AverageSpeed)
					lr.MaxSpeed = float64(act.MaxSpeed)
					lr.SportType = string(act.SportType)
				}
			}
			param := sqlite.CreateActivityParams{
				ID:           act.ID,
				UserID:       act.Athlete.ID,
				CreateAt:     currentTime,
				StartDate:    startDate,
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
		//lParams := sqlite.UpdateLongestRunPerDayParams{
		//	UserID:       lr.UserID,
		//	Today:        lr.Today,
		//	ActivityID:   lr.ActivityID,
		//	StartDate:    lr.StartDate,
		//	Distance:     lr.Distance,
		//	AverageSpeed: lr.AverageSpeed,
		//	MovingTime:   lr.MovingTime,
		//	Name:         lr.Name,
		//	SportType:    lr.SportType,
		//	MaxSpeed:     lr.MaxSpeed,
		//}
		if err = s.db.CreateActivityTx(ctx, params); err != nil {
			return fmt.Errorf("CreateActivityTx failed: %s", err)
		}
	}
	return nil
}
