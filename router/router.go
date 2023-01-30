package router

import (
	"WakaTImeGo/router/userRouter"
	"WakaTImeGo/router/wakatime/v1"
	"github.com/gin-gonic/gin"
)

func InitRotesNeedAuth(engine *gin.Engine) {
	//init userController router
	userRouter.InitUserRoutesNeedAuth(engine)
	v1.InitTodayRoutes(engine)
	v1.InitHeartbeatRoutes(engine)
	v1.InitDurationRoutes(engine)
}

func InitRotesNotNeedAuth(engine *gin.Engine) {
	userRouter.InitUserRoutesNotNeedAuth(engine)
}
