package router

import (
	"WakaTImeGo/router/userRouter"
	"WakaTImeGo/router/wakatime/v1/today"
	"github.com/gin-gonic/gin"
)

func InitRotesNeedAuth(engine *gin.Engine) {
	//init userController router
	userRouter.InitUserRoutes()
	today.InitTodayRoutes(engine)
}
