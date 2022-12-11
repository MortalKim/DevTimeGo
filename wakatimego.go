package main

import (
	"WakaTImeGo/basic/authentication"
	"WakaTImeGo/basic/database"
	"WakaTImeGo/config"
	"WakaTImeGo/user/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var log = logrus.New()

func main() {
	config.InitConfig()
	database.InitDatabase()
	service.InitDatabase()
	initRoute()
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
	r.GET("/service_with_auth", authentication.ServiceWithAuth)
	//#endregion Some api need authentication

	r.Run()
}
