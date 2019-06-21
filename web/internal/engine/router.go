package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/m-go-cookbook/web/internal/ctrl"
)

func Default() *gin.Engine {
	engine := gin.New()
	if gin.Mode() != gin.ReleaseMode {
		engine.Use(gin.Logger())
	}
	engine.Use(recovery)
	engine.Use(responseHandler)

	engine.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	apiGroup := engine.Group("/api")
	// ------------ example ----------------- //
	apiGroup.POST("/examples", ctrl.Example.Create)
	apiGroup.GET("/examples", ctrl.Example.List)
	apiGroup.GET("/examples/empty", ctrl.Example.EmptyBody)
	apiGroup.GET("/examples/404", ctrl.Example.With404)
	// ------------ example ----------------- //

	return engine
}
