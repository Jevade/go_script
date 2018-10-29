package router

import (
	"net/http"

	"../handler/sd"
	"../handler/user"
	"./middleware"
	"github.com/gin-gonic/gin"
)

//Load load middlerware,rouer,handles
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Option)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/v1/user")
	{
		u.POST("/:username", user.Create)
		u.GET("/:username", user.Info)
	}

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/disk", sd.DiskCheck)
	}

	return g

}
