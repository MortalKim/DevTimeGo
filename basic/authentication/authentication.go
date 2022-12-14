package authentication

import (
	"WakaTImeGo/config"
	"WakaTImeGo/constant"
	"WakaTImeGo/model/entity"
	"WakaTImeGo/service/duration"
	"WakaTImeGo/service/userService"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"strings"
)

// Authentication is a middleware that checks if the request is authenticated.
// If the request is authenticated, the middleware will call the next handler
// to process the request.
// If the request is not authenticated, the middleware will abort the request
// and return a 401 status code.

const (
	TokenSalt = "default_salt"
)

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		//judge token prefix is "basic"
		if strings.HasPrefix(token, "Basic ") {
			//get string after "basic "
			token = strings.TrimPrefix(token, "Basic ")
			decodeToken, _ := base64.StdEncoding.DecodeString(token)
			//Get userService by token
			user, err := userService.GetUserByApiKey(string(decodeToken))

			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Request.Header.Add(constant.DECRYPTED_USER_ID, strconv.Itoa(int(user.ID)))
			c.Next()
		} else if strings.HasPrefix(token, "Bearer ") {
			token = strings.TrimPrefix(token, "Bearer ")
			jwtStruct := entity.JwtStruct{}
			token, err := jwt.ParseWithClaims(token, &jwtStruct, func(token *jwt.Token) (i interface{}, err error) {
				return []byte(config.JWT_SECRET), nil
			})
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Request.Header.Add(constant.DECRYPTED_USER_ID, strconv.Itoa(int(jwtStruct.UserID)))
			if _, ok := token.Claims.(*entity.JwtStruct); ok && token.Valid {
				c.Next()
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//if strings.ToLower(MD5([]byte(username+ts+TokenSalt))) == strings.ToLower(token) {
		//	c.Next()
		//} else {
		//	c.Abort()
		//	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		//	return
		//}
	}
}

func ServiceWithoutAuth(c *gin.Context) {
	duration.GenericUserDuration("1")
}

func ServiceWithAuth(c *gin.Context) {

}
