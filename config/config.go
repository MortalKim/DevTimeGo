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
