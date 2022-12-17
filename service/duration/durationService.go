package duration

import (
	"WakaTImeGo/basic/database"
	"WakaTImeGo/model/entity"
	"WakaTImeGo/service/heartbeat"
	log "github.com/sirupsen/logrus"
	"time"
)

/**
 * @Author: Kim
 * @Description: to generic user's coding activity for the given day as an array of durations.
 * @File:  durationService
 * @Date: 12/13/2022 6:47 PM
 */

func InitDatabase() {
	//create table to database
	err := database.GetDb().AutoMigrate(&entity.Duration{})
	if err != nil {
		return
	}
}

func GenericUserDuration(userID string) {
	//Get user's mHeartbeat not counted, sort by time
	heartbeats := heartbeat.GetHeartbeatsNotCountedByUser(userID)
	//judge heartbeats is empty
	if len(heartbeats) == 0 {
		return
	}
	//Get user's last Duration by heartbeats time
	lastDuration, isValid := getLastDurationOfTime(userID, heartbeats[0].Time)
	//Generic user's duration
	for _, mHeartbeat := range heartbeats {
		//if mHeartbeat is not in last duration, or it has diff content, then create a new duration
		if isValid && mHeartbeat.Time.Sub(lastDuration.Time) < time.Minute*5 && mHeartbeat.Project == lastDuration.Project &&
			mHeartbeat.Language == lastDuration.Language && mHeartbeat.Editor == lastDuration.Editor &&
			mHeartbeat.OperatingSystem == lastDuration.OperatingSystem &&
			mHeartbeat.Machine == lastDuration.Machine &&
			mHeartbeat.Branch == lastDuration.Branch {
			//if mHeartbeat is in last duration, then update last duration
			//Duration is the sum of all heartbeats in it. so it is the last mHeartbeat's time minus the first mHeartbeat's time.
			//To make last mHeartbeat have effect, so add 2 minutes to it.
			lastDuration.Duration = mHeartbeat.Time.Sub(lastDuration.Time) + 2*time.Minute
			lastDuration.NumHeartbeats++
			err := lastDuration.Update()
			if err != nil {
				log.Error(err)
				return
			}
		} else {
			//For new duration, there are 2 minutes for a single mHeartbeat of it.
			duration := entity.Duration{
				UserID:          userID,
				Time:            mHeartbeat.Time,
				Duration:        2 * time.Minute,
				Project:         mHeartbeat.Project,
				Language:        mHeartbeat.Language,
				Editor:          mHeartbeat.Editor,
				OperatingSystem: mHeartbeat.OperatingSystem,
				Machine:         mHeartbeat.Machine,
				Branch:          mHeartbeat.Branch,
				NumHeartbeats:   1,
			}
			err := duration.Add()
			if err != nil {
				log.Error(err)
				return
			}
			lastDuration = duration
			isValid = true
		}
		mHeartbeat.IsCounted = true
		err := mHeartbeat.Update()
		if err != nil {
			log.Error(err)
			return
		}
	}
}

// get user's last duration
func getLastDuration(userID string) entity.Duration {
	var duration entity.Duration
	database.GetDb().Where("user_id = ?", userID).Last(&duration)
	return duration
}

func getLastDurationOfTime(userID string, time time.Time) (entity.Duration, bool) {
	var duration entity.Duration
	err := database.GetDb().Where("user_id = ? AND time < ?", userID, time).Last(&duration).Error
	if err != nil {
		return entity.Duration{}, false
	}
	return duration, true
}
