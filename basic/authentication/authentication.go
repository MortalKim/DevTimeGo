package authentication

import (
	"WakaTImeGo/constant"
	"WakaTImeGo/service/userService"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
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
		if strings.HasPrefix(token, "Basic") {
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

}

func ServiceWithAuth(c *gin.Context) {

}
