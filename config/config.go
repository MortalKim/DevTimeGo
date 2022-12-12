package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	runMode := viper.GetString("runmode")
	log.Info("Run mode: ", runMode)
}

func GetRedisConfig() (host, password string, port, db int) {
	host = viper.GetString("redis.host")
	port = viper.GetInt("redis.port")
	password = viper.GetString("redis.password")
	db = viper.GetInt("redis.db")
	return host, password, port, db
}

func GetDatabaseConfig() (host, username, password, dbname string, port string) {
	host = viper.GetString("database.host")
	port = viper.GetString("database.port")
	username = viper.GetString("database.username")
	password = viper.GetString("database.password")
	dbname = viper.GetString("database.dbname")
	return host, username, password, dbname, port
}
