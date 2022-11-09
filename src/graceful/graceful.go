package graceful

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func Graceful(dir string, hot bool) {
	thread := make(chan os.Signal)
	signal.Notify(thread, os.Interrupt)
	go func() {
		<-thread // wait for os interrupt
		fmt.Println("\nGracefully shutting down")
		// if watcher.Watch != nil {
		// 	watcher.Watch.Close()
		// }
		reader, err := os.ReadFile(dir + "/index.html")
		if err != nil {
			panic(err)
		}
		var data string
		if hot {
			data = strings.Replace(string(reader), "<script src=\"/blast-ws.js\" type=\"module\" defer></script>", "{{.Inject}}", 1)
		} else {
			data = strings.Replace(string(reader), "<!-- {{.Inject}} -->", "{{.Inject}}", 1)
		}
		os.WriteFile(dir+"/index.html", []byte(data), 0644)
		os.Remove(dir + "/blast-ws.js")
		os.Exit(0)
	}()
}
