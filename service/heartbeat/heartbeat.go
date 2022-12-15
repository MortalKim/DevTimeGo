package heartbeat

import (
	"WakaTImeGo/basic/database"
	"WakaTImeGo/model/entity"
)

func InitDatabase() {
	//create table to database
	err := database.GetDb().AutoMigrate(&entity.Heartbeat{})
	if err != nil {
		return
	}
}

// GetHeartbeatByTime Get heartbeat form start to end time
func GetHeartbeatByTime(startTime, endTime string) []entity.Heartbeat {
	var heartbeats []entity.Heartbeat
	//sort by time
	database.GetDb().Where("time >= ? AND time <= ?", startTime, endTime).Order("time").Find(&heartbeats)

	return heartbeats
}
