package duration

import (
	"WakaTImeGo/basic/database"
	"WakaTImeGo/model/entity"
	"WakaTImeGo/model/entity/request/duration"
	"WakaTImeGo/service/heartbeat"
	"WakaTImeGo/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
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
	//Get user's Durations by heartbeats time
	//Get duration of first heartbeat time's day start time, to last heartbeat time's day end time
	startTime, _ := utils.GetDayTime(heartbeats[0].Time)
	_, endTime := utils.GetDayTime(heartbeats[len(heartbeats)-1].Time)
	durations, isEmpty := GetDurationByTime(userID, startTime, endTime)

	//When durations updated, maybe some duration has out of date, so we need to delete it
	toDeleteDurations := make([]entity.Duration, 0)
	var isValid = !isEmpty
	var lastDurationIndex int = 0
	//start at first
	//Generic user's duration
	for index, mHeartbeat := range heartbeats {
		log.Info("Process heartbeat to duration, total" + strconv.Itoa(len(heartbeats)) + ", current: " + strconv.Itoa(index))
		var lastDuration entity.Duration
		lastDurationIndex = 0
		//find current heartbeat's last duration
		for lastDurationIndex < len(durations) {
			if durations[lastDurationIndex].Time.Before(mHeartbeat.Time) {
				if lastDurationIndex < len(durations)-1 {
					if durations[lastDurationIndex+1].Time.After(mHeartbeat.Time) {
						//it means current heartbeat is in this two durations
						break
					} else {
						//it means current heartbeat is not in this two durations
						lastDurationIndex++
						continue
					}
				} else {
					//if it is durations last element, it is last duration
					break
				}
			}
			lastDurationIndex++
		}
		if lastDurationIndex >= len(durations) {
			isValid = false
		} else {
			lastDuration = durations[lastDurationIndex]
		}

		//1. If mHeartbeat's time to last duration's time is less than 5 minutes, and it has same content, then update last duration
		//1.1. If last duration is reach the next duration after update, and it has same content, merge it
		//1.2. If last duration is reach the next duration and it has different content, update duration time to next duration's start time
		//2. If mHeartbeat's time to last duration's time is more than 5 minutes, or diff content, then create a new duration, call it last duration
		//2.1 If new last duration is reach the next duration after update, and it has same content, merge it
		//2.2 If new last duration is reach the next duration and it has different content, update duration time to next duration's start time
		//3. If mHeartbeat's time to last duration's time is more than 5 minutes, and reach the next duration, last duration = next duration
		if isValid && mHeartbeat.Time.Sub(lastDuration.Time) < time.Minute*5 && mHeartbeat.Project == lastDuration.Project &&
			mHeartbeat.Language == lastDuration.Language && mHeartbeat.Editor == lastDuration.Editor &&
			mHeartbeat.OperatingSystem == lastDuration.OperatingSystem &&
			mHeartbeat.Machine == lastDuration.Machine &&
			mHeartbeat.Branch == lastDuration.Branch &&
			mHeartbeat.Category == lastDuration.Category {
			//if mHeartbeat is in last duration, then update last duration
			//Duration is the sum of all heartbeats in it. so it is the last mHeartbeat's time minus the first mHeartbeat's time.
			//To make last mHeartbeat have effect, so add 2 minutes to it.
			lastDuration.Duration = mHeartbeat.Time.Sub(lastDuration.Time) + 2*time.Minute
			lastDuration.NumHeartbeats++
			if lastDurationIndex+1 < len(durations) && lastDuration.Time.Add(lastDuration.Duration).After(durations[lastDurationIndex+1].Time) {
				//if last duration is reach the next duration after update, and it has same content, merge it
				if mHeartbeat.Project == durations[lastDurationIndex+1].Project &&
					mHeartbeat.Language == durations[lastDurationIndex+1].Language && mHeartbeat.Editor == durations[lastDurationIndex+1].Editor &&
					mHeartbeat.OperatingSystem == durations[lastDurationIndex+1].OperatingSystem &&
					mHeartbeat.Machine == durations[lastDurationIndex+1].Machine &&
					mHeartbeat.Branch == durations[lastDurationIndex+1].Branch &&
					mHeartbeat.Category == durations[lastDurationIndex+1].Category {

					lastDuration.Duration = durations[lastDurationIndex+1].Time.Add(durations[lastDurationIndex+1].Duration).Sub(lastDuration.Time)
					lastDuration.NumHeartbeats = lastDuration.NumHeartbeats + durations[lastDurationIndex+1].NumHeartbeats
					toDeleteDurations = append(toDeleteDurations, durations[lastDurationIndex+1])
					//lastDurationIndex++
				} else {
					//if last duration is reach the next duration and it has different content, update duration time to next duration's start time
					lastDuration.Duration = durations[lastDurationIndex+1].Time.Sub(lastDuration.Time)
				}
			}
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
				Category:        mHeartbeat.Category,
				Project:         mHeartbeat.Project,
				Language:        mHeartbeat.Language,
				Editor:          mHeartbeat.Editor,
				OperatingSystem: mHeartbeat.OperatingSystem,
				Machine:         mHeartbeat.Machine,
				Branch:          mHeartbeat.Branch,
				NumHeartbeats:   1,
			}
			needAddToDatabase := true
			if lastDurationIndex+1 < len(durations) && duration.Time.Add(duration.Duration).After(durations[lastDurationIndex+1].Time) {
				//if new last duration is reach the next duration after update, and it has same content, merge it
				if mHeartbeat.Project == durations[lastDurationIndex+1].Project &&
					mHeartbeat.Language == durations[lastDurationIndex+1].Language && mHeartbeat.Editor == durations[lastDurationIndex+1].Editor &&
					mHeartbeat.OperatingSystem == durations[lastDurationIndex+1].OperatingSystem &&
					mHeartbeat.Machine == durations[lastDurationIndex+1].Machine &&
					mHeartbeat.Branch == durations[lastDurationIndex+1].Branch &&
					mHeartbeat.Category == durations[lastDurationIndex+1].Category {

					duration.Duration = durations[lastDurationIndex+1].Time.Add(durations[lastDurationIndex+1].Duration).Sub(duration.Time)
					duration.NumHeartbeats = duration.NumHeartbeats + durations[lastDurationIndex+1].NumHeartbeats
					//delete unnecessary duration
					toDeleteDurations = append(toDeleteDurations, durations[lastDurationIndex+1])
					//lastDurationIndex++
				} else {
					//if new last duration is reach the next duration and it has different content, update duration time to next duration's start time
					duration.Duration = durations[lastDurationIndex+1].Time.Sub(duration.Time)
				}
			}
			if lastDuration.Time.Add(lastDuration.Duration).After(duration.Time) {
				//if last duration is reach the next duration, last duration = next duration
				lastDuration.Duration = duration.Time.Sub(lastDuration.Time)
				needAddToDatabase = false
				err := lastDuration.Update()
				if err != nil {
					log.Error(err)
				}
			}
			if needAddToDatabase {
				err := duration.Add()
				if err != nil {
					log.Error(err)
				}
				lastDuration = duration

				if lastDurationIndex < len(durations) {
					durations = append(durations[:lastDurationIndex+1], append([]entity.Duration{lastDuration}, durations[lastDurationIndex+1:]...)...)
					// insert lastDuration to durations by before durations[lastDurationIndex}
					//durations = append(durations, entity.Duration{})
					//copy(durations[lastDurationIndex+1:], durations[lastDurationIndex:])
					//durations[lastDurationIndex] = lastDuration
				} else {
					durations = append(durations, lastDuration)
				}
			}

			isValid = true
		}
		mHeartbeat.IsCounted = true
		mHeartbeat.DurationId = lastDuration.ID
		err := mHeartbeat.Update()
		if err != nil {
			log.Error(err)
		}
	}
	for _, duration := range toDeleteDurations {
		err := duration.Delete()
		if err != nil {
			log.Error(err)
		}
	}
	log.Info("Generic user's duration success")
}

// GetDurationByTime get user's durations by time
// param: userID, startTime, endTime
// return: durations, isEmpty
func GetDurationByTime(userID string, startTime, endTime time.Time) ([]entity.Duration, bool) {
	var durations []entity.Duration
	//get and sort by time asc
	database.GetDb().Where("user_id = ? AND time >= ? AND time <= ?", userID, startTime, endTime).Order("time asc").Find(&durations)
	if len(durations) == 0 {
		return durations, true
	}
	return durations, false
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

func GetDurationByParams(params duration.SearchParams) []entity.Duration {
	var durations []entity.Duration
	timeStr := time.UnixMilli(params.Date)
	startDate, endDate := utils.GetDayTime(timeStr)
	database.GetDb().Where("user_id = ? AND time >= ? AND time <= ?",
		params.UserID, startDate, endDate).Find(&durations)
	return durations
}
