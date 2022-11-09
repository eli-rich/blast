package watcher

import (
	"log"

	"github.com/rjeczalik/notify"
)

var watcherChannel chan notify.EventInfo

func Init(hot bool) chan notify.EventInfo {
	if !hot {
		return nil
	}
	watcherChannel = make(chan notify.EventInfo, 1)
	return watcherChannel
}

func Watch(dir string, thread chan notify.EventInfo) {
	if err := notify.Watch(dir, watcherChannel, notify.All); err != nil {
		log.Fatalln(err)
	}
	defer notify.Stop(watcherChannel)
	for {
		thread <- <-watcherChannel
	}
}
