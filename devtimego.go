package main

import (
	"WakaTImeGo/basic/authentication"
	"WakaTImeGo/basic/database"
	"WakaTImeGo/basic/json"
	"WakaTImeGo/basic/redis"
	"WakaTImeGo/config"
	"WakaTImeGo/router"
	"WakaTImeGo/service/duration"
	"WakaTImeGo/service/heartbeat"
	"WakaTImeGo/service/userService"
	"WakaTImeGo/task"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func main() {
	config.InitConfig()
	initCustomJsonDecoder()
	initDatabase()
	redis.Setup()
	go task.InitTaskService()
	initRoute()
}

func initCustomJsonDecoder() {
	json.RegisterTimeDecoderFunc()
}

func initDatabase() {
	database.InitDatabase()
	userService.InitDatabase()
	heartbeat.InitDatabase()
	database.InitDatabase()
	duration.InitDatabase()
}

func initRoute() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//#region Some api do not need authentication
	r.GET("/service_without_auth", authentication.ServiceWithoutAuth)
	//#endregion Some api do not need authentication

	//#region Some api need authentication
	r.Use(authentication.Authorize())

	router.InitRotesNeedAuth(r)

	//r.GET("/service_with_auth", authentication.ServiceWithAuth)
	//#endregion Some api need authentication

	r.Run()
}
