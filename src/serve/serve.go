package serve

import (
	"fmt"
	"time"

	"github.com/BlazingFire007/blast/src/logger"
	"github.com/gofiber/fiber/v2"
)

func Serve(port int, dir string) {
	app := fiber.New()
	app.Static("/", string(dir), fiber.Static{
		CacheDuration: 100 * time.Millisecond,
		ModifyResponse: func(c *fiber.Ctx) error {
			c.Response().Header.Add("Cache-Control", "no-store")
			return nil
		},
	})

	logger.Success(fmt.Sprintf("Serving %s", dir))
	app.Listen(fmt.Sprintf(":%d", port))
}
