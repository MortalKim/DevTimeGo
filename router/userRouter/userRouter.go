package userRouter

import "github.com/gin-gonic/gin"
import "WakaTImeGo/controller/userController"

func InitUserRoutesNeedAuth(r *gin.Engine) {
	//init userController router
	user := r.Group("/user")
	{
		user.GET("/info", userController.GetUserInfo)
		user.PUT("/info", userController.UpdateUserInfo)
		user.PUT("/password", userController.UpdatePassword)
	}
}

func InitUserRoutesNotNeedAuth(r *gin.Engine) {
	//init userController router
	user := r.Group("/user")
	{
		user.POST("/login", userController.Login)
		user.POST("/register", userController.Register)
	}
}
