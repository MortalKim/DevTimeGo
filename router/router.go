package router

import (
	"WakaTImeGo/router/userRouter"
	"WakaTImeGo/router/wakatime/v1/heartbeat"
	"WakaTImeGo/router/wakatime/v1/today"
	"github.com/gin-gonic/gin"
)

func InitRotesNeedAuth(engine *gin.Engine) {
	//init userController router
	userRouter.InitUserRoutesNeedAuth(engine)
	today.InitTodayRoutes(engine)
	heartbeat.InitHeartbeatRoutes(engine)
}

func InitRotesNotNeedAuth(engine *gin.Engine) {
	userRouter.InitUserRoutesNotNeedAuth(engine)
}
