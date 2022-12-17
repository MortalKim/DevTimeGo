package today

import (
	"WakaTImeGo/basic/authentication"
	"WakaTImeGo/constant"
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
	today := engine.Group(constant.API_V1_PREFIX)
	{
		today.GET("/users/:user/statusbar/today", todayController.GetToday)
	}
}
