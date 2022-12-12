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
