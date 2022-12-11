package authentication

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
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
		username := c.Query("username")
		ts := c.Query("ts")
		token := c.Request.Header.Get("Authorization")

		if strings.ToLower(MD5([]byte(username+ts+TokenSalt))) == strings.ToLower(token) {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "访问未授权"})
			return
		}
	}
}

func ServiceWithoutAuth(c *gin.Context) {

}

func ServiceWithAuth(c *gin.Context) {

}
