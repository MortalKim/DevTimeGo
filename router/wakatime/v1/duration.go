package v1

import (
	"WakaTImeGo/constant"
	v1 "WakaTImeGo/controller/wakatime/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Kim
 * @Description:
 * @File:  duration
 * @Date: 12/19/2022 10:36 AM
 */

func InitDurationRoutes(engine *gin.Engine) {
	//init userController router
	duration := engine.Group(constant.API_V1_PREFIX)
	{
		duration.GET("/users/:user/durations", v1.GetDurationsByParams)
		duration.GET("/users/current/durations", v1.GetDurationsByParams)
	}
}
