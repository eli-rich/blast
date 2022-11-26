package serve

import (
	"fmt"

	"github.com/BlazingFire007/blast/src/logger"
	"github.com/gin-gonic/gin"
)

func Serve(port int, dir string) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Static("/", string(dir))

	logger.Success(fmt.Sprintf("Serving %s", dir))
	app.Run(fmt.Sprintf(":%d", port))
}
