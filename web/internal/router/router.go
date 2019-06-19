package router

import (
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/iceyang/m-go-cookbook/web/internal"
	"github.com/iceyang/m-go-cookbook/web/internal/ctrl"
)

func responseHandler(c *gin.Context) {
	c.Next()
	if c.Writer.Status() == http.StatusNotFound && c.Writer.Size() <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}
	if c.Writer.Status() == http.StatusOK && c.Writer.Size() <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	}
}

func recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 65536)
			buf = buf[:runtime.Stack(buf, false)]
			e, ok := err.(*internal.Error)
			if ok {
				if e.Code >= 500 {
					log.Printf("%s\n%s", err, buf)
				}
				c.AbortWithStatusJSON(e.Code, gin.H{
					"message": e.Msg,
				})
				return
			}
			log.Printf("%s\n%s", err, buf)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "InternalServerError",
			})
		}
	}()
	c.Next()
}

func Default() *gin.Engine {
	router := gin.New()
	if gin.Mode() != gin.ReleaseMode {
		router.Use(gin.Logger())
	}
	router.Use(recovery)
	router.Use(responseHandler)

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	apiGroup := router.Group("/api")
	// ------------ example ----------------- //
	apiGroup.POST("/examples", ctrl.Example.Create)
	apiGroup.GET("/examples", ctrl.Example.List)
	apiGroup.GET("/examples/empty", ctrl.Example.EmptyBody)
	apiGroup.GET("/examples/404", ctrl.Example.With404)
	// ------------ example ----------------- //

	return router
}
