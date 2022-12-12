package today

import (
	"WakaTImeGo/basic/authentication"
	todayController "WakaTImeGo/controller/wakatime/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Kim
 * @Description: Init wakatime today status bar router
 * @File:  today
 * @Date: 12/12/2022 3:45 PM
 */

func InitTodayRoutes(engine *gin.Engine) {
	engine.GET("/service_with_auth", authentication.ServiceWithAuth)
	//init userController router
	today := engine.Group("/api/v1/users")
	{
		today.GET("/:user/statusbar/today", todayController.GetToday)
	}
}
