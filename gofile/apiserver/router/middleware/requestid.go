package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/satori/go.uuid"
)

//RequestID will set "X-Request-Id" with uuid in head
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-Id")
		if requestID == "" {
			u4, _ := uuid.NewV4()
			requestID = u4.String()
		}
		c.Set("X-Request-Id", requestID)
		// rid, _ := c.Get("X-Request-Id")
		log.Warn(requestID)
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
