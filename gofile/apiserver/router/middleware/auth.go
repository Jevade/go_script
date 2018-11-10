package middleware

import (
	"../../handler"
	"../../pkg/errno"
	"../../pkg/token"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware is to authritify the token and secret
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
