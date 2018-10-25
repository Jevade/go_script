package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//NoCache prevent caching http responses
func NoCache(c *gin.Context) {
	c.Header("Cacha-Control", "no-cache,no-store,max-age=0,must-revalidate,value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

//Option deal with option request
func Option(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS,PATCH,DELETE")
		c.Header("Access-Control-Allow-Headers", "authorization,origin,content-type,accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,OPTIONS,PATCH,DELETE")
		c.Header("Content-Cype", "application/json")
		c.AbortWithStatus(200)
	}

}

//Secure add security headers
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1;mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}

}
