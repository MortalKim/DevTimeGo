package router

import "WakaTImeGo/router/userRouter"

func InitRotes() {
	//init userController router
	userRouter.InitUserRoutes()
}
