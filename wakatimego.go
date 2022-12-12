package main

import (
	"WakaTImeGo/basic/authentication"
	"WakaTImeGo/basic/database"
	"WakaTImeGo/basic/redis"
	"WakaTImeGo/config"
	"WakaTImeGo/router"
	"WakaTImeGo/service/heartbeat"
	"WakaTImeGo/service/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func main() {
	config.InitConfig()
	initDatabase()
	redis.Setup()
	initRoute()
}

func initDatabase() {
	database.InitDatabase()
	user.InitDatabase()
	heartbeat.InitDatabase()
}

func initRoute() {
	r := gin.Default()

	router.InitRotes()

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
	r.GET("/service_with_auth", authentication.ServiceWithAuth)
	//#endregion Some api need authentication

	r.Run()
}
