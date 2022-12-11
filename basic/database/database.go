package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB

func InitDatabase() {
	//Connect to mysql
	dsn := "root:root@tcp(192.168.1.19:3306)/wakatimekt?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	instance = db
	if err != nil {
		panic(err)
	}
}

func GetDb() *gorm.DB {
	return instance
}
