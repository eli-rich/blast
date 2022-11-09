package serve

import (
	"fmt"
	"time"

	"github.com/BlazingFire007/blast/src/logger"
	"github.com/BlazingFire007/blast/src/watcher"
	"github.com/gin-gonic/gin"
)

var HOT bool

func Serve(port int, dir string, hot bool) {
	HOT = hot
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	Initialize(router, dir)
	AddRoutes(dir)

	if HOT {
		router.GET("/blast/ws", func(c *gin.Context) {
			wshandler(c.Writer, c.Request)
		})
		time.Sleep(500 * time.Millisecond)
		watcherChannel := watcher.Init(HOT)
		go watcher.Watch(dir, watcherChannel)
	}

	logger.Success(fmt.Sprintf("Serving %s on port %d", dir, port))
	router.Run(fmt.Sprintf(":%d", port))
}
