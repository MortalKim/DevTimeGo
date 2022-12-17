package heartbeat

import (
	"WakaTImeGo/constant"
	"WakaTImeGo/controller/wakatime/v1"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: Kim
 * @Description: wakatime v1 api heartbeat
 * @File:  heartbeat
 * @Date: 12/12/2022 5:00 PM
 */

func InitHeartbeatRoutes(engine *gin.Engine) {
	//init userController router
	heartbeat := engine.Group(constant.API_V1_PREFIX)
	{
		heartbeat.POST("/users/:user/heartbeats.bulk", v1.SaveHeartbeats)
		heartbeat.POST("/users/:user/heartbeats", v1.SaveHeartbeat)
	}
}
