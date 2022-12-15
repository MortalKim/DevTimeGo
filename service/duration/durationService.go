package duration

import (
	"WakaTImeGo/basic/database"
	"WakaTImeGo/model/entity"
	"WakaTImeGo/service/heartbeat"
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
	//Get user's last Duration
	lastDuration := getLastDuration(userID)
	//Get user's heartbeat from last duration to now, sort by time
	heartbeats := heartbeat.GetHeartbeatByTime(lastDuration.Time.Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	//Generic user's duration
	for _, heartbeat := range heartbeats {
		//if heartbeat is not in last duration, then create a new duration
		if heartbeat.Time.Sub(lastDuration.Time) > time.Minute*5 {
			//For new duration, there are 2 minutes for a single heartbeat of it.
			duration := entity.Duration{
				UserID:          userID,
				Time:            heartbeat.Time,
				Duration:        2 * time.Minute,
				Project:         heartbeat.Project,
				Language:        heartbeat.Language,
				Editor:          heartbeat.Editor,
				OperatingSystem: heartbeat.OperatingSystem,
				Machine:         heartbeat.Machine,
				Branch:          heartbeat.Branch,
				NumHeartbeats:   1,
			}
			duration.Add()
			lastDuration = duration
		} else {
			//if heartbeat is in last duration, then update last duration
			//Duration is the sum of all heartbeats in it. so it is the last heartbeat's time minus the first heartbeat's time.
			//To make last heartbeat have effect, so add 2 minutes to it.
			lastDuration.Duration = heartbeat.Time.Sub(lastDuration.Time) + 2*time.Minute
			lastDuration.NumHeartbeats++
			lastDuration.Update()
		}
	}
}

//get user's last duration
func getLastDuration(userID string) entity.Duration {
	var duration entity.Duration
	database.GetDb().Where("user_id = ?", userID).Last(&duration)
	return duration
}
