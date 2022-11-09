package serve

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "hello" {
			conn.WriteMessage(websocket.TextMessage, []byte("hello"))
			go pingInterval(conn)
		}
	}
}

func pingInterval(conn *websocket.Conn) {
	for {
		time.Sleep(10 * time.Second)
		conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	}
}
