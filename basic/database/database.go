package database

import (
	"WakaTImeGo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB

func InitDatabase() {
	//Connect to mysql
	host, username, password, dbname, port := config.GetDatabaseConfig()
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	instance = db
	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	return instance
}
