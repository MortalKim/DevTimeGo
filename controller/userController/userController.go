package userController

import (
	"WakaTImeGo/config"
	"WakaTImeGo/constant"
	"WakaTImeGo/model/entity"
	"WakaTImeGo/model/entity/request"
	"WakaTImeGo/model/entity/response"
	"WakaTImeGo/model/response/base"
	"WakaTImeGo/service/userService"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func Auth(c *gin.Context) {
	//get user info by token
	userId := c.Request.Header.Get(constant.DECRYPTED_USER_ID)
	user, err := userService.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	userInfo := make(map[string]interface{})
	userInfo["username"] = user.UserName
	userInfo["email"] = user.Email
	userInfo["userID"] = user.ID
	c.JSON(http.StatusOK, userInfo)
}

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
	jwtStruct := entity.JwtStruct{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24 * 7),
			},
			Issuer: "DevTimeGO",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtStruct)
	signedString, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	res.Code = http.StatusOK
	res.Message = "login success"
	res.Data = response.LoginResponse{
		Token: signedString,
	}
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
