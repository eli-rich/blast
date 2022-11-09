package serve

import (
	"fmt"

	"github.com/BlazingFire007/blast/src/logger"
	"github.com/gin-gonic/gin"
)

var HOT bool

func Serve(port int, dir string, hot bool) {
	HOT = hot
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/blast/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
	createRoutes(router, dir)
	logger.Success(fmt.Sprintf("Serving %s on port %d", dir, port))
	router.Run(fmt.Sprintf(":%d", port))
}
