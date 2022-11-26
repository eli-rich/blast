package serve

import (
	"fmt"

	"github.com/BlazingFire007/blast/src/logger"
	"github.com/gin-gonic/gin"
)

func DisableCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Header("Server", "Blast")
		c.Next()
	}
}

func Serve(port int, dir string) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(DisableCache())
	app.Static("/", string(dir))

	logger.Success(fmt.Sprintf("Serving %s", dir))
	app.Run(fmt.Sprintf(":%d", port))
}
