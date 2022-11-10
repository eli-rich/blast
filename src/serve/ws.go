package serve

import (
	"log"
	"time"

	"github.com/BlazingFire007/blast/src/watcher"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var PING_INTERVAL = 1 * time.Second
var LAST_RESPONSE = time.Now()

func wsmiddle(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func wsHandler(c *websocket.Conn) {

	err := c.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		log.Fatalln(err)
	}
	listenForReload(c)

}
func listenForReload(c *websocket.Conn) {
	change := <-watcher.WatcherChannel
	log.Println(change)
	c.WriteMessage(websocket.TextMessage, []byte("reload"))
	c.Close()
}
