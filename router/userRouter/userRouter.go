package userRouter

import "github.com/gin-gonic/gin"
import "WakaTImeGo/controller/userController"

func InitUserRoutes() {
	router := gin.Default()
	//init userController router
	user := router.Group("/userController")
	{
		user.POST("/login", userController.Login)
		user.POST("/register", userController.Register)
		user.GET("/info", userController.GetUserInfo)
		user.PUT("/info", userController.UpdateUserInfo)
		user.PUT("/password", userController.UpdatePassword)
	}
}
