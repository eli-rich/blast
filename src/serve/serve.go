package serve

import (
	"fmt"
	"time"

	"github.com/BlazingFire007/blast/src/logger"
	"github.com/BlazingFire007/blast/src/watcher"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var HOT bool

func Serve(port int, dir string, hot bool) {
	HOT = hot
	app := fiber.New()
	initializeRouter(dir, app)
	initializePaths(dir)
	app.Static("/", string(ROOT), fiber.Static{
		CacheDuration: 100 * time.Millisecond,
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Response().Header.Add("Cache-Control", "no-store")
			return nil
		},
	})
	// app.Use(cache.Config{

	// })
	if HOT {
		app.Use("/ws", wsmiddle)
		app.Get("/ws:blast", websocket.New(wsHandler))
		time.Sleep(500 * time.Millisecond)
		watcher.Init(HOT)
		go watcher.Watch(dir, watcher.WatcherChannel)
	}
	logger.Success(fmt.Sprintf("Serving %s", dir))
	app.Listen(fmt.Sprintf(":%d", port))
}
