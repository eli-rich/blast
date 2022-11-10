package watcher

import (
	"log"

	"github.com/rjeczalik/notify"
)

var WatcherChannel chan notify.EventInfo

func Init(hot bool) chan notify.EventInfo {
	if !hot {
		return nil
	}
	WatcherChannel = make(chan notify.EventInfo, 1)
	return WatcherChannel
}

func Watch(dir string, thread chan notify.EventInfo) {
	if err := notify.Watch(dir, WatcherChannel, notify.All); err != nil {
		log.Fatalln(err)
	}
	defer notify.Stop(WatcherChannel)
	for {
		thread <- <-WatcherChannel
	}
}
