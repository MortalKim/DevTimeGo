package userController

import (
	"WakaTImeGo/model/entity/request"
	"WakaTImeGo/model/response/base"
	"WakaTImeGo/service/userService"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var userReq request.UserRegister

	err := c.BindJSON(&userReq)
	res := base.ResponseBase{}
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "Json Error"
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	user, err := userService.GetUserByEmail(userReq.Email)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "user not exist"
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Code = http.StatusOK
	res.Message = "login success"
	res.Data = user
	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {
	var userReq request.UserRegister

	err := c.BindJSON(&userReq)
	res := base.ResponseBase{}
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "Json Error"
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	canRegister, info := userService.CanRegister(userReq.Username, userReq.Email)

	if !canRegister {
		res.Code = http.StatusInternalServerError
		res.Message = info
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	_, err = userService.Register(userReq.Username, userReq.Password, userReq.Email)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "register failed"
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Code = http.StatusOK
	res.Message = "register success"
	c.JSON(http.StatusOK, res)
}

func GetUserInfo(c *gin.Context) {

}

func UpdateUserInfo(c *gin.Context) {

}

func UpdatePassword(c *gin.Context) {

}
