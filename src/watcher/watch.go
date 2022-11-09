package watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

var Watch *fsnotify.Watcher

func RootWatcher(hot bool, dir string, thread chan bool) {
	if !hot {
		return
	}
	watcher, err := fsnotify.NewWatcher()
	Watch = watcher
	if err != nil {
		log.Fatalln(err)
	}
	defer Watch.Close()
	err = Watch.Add(dir)
	if err != nil {
		log.Fatalln(err)
	}
	go listen(thread)
	<-make(chan struct{})
}

func listen(thread chan bool) {
	for {
		select {
		case event, ok := <-Watch.Events:
			if !ok {
				return
			}
			// log.Println("event:", event)
			if event.Has(fsnotify.Write) {
				thread <- true
			}
			// Code below is buggy, TODO: fix. Some events fire twice and remove does not fire on macOS.
			/* else if event.Has(fsnotify.Create) {
				thread <- true
			} else if event.Has(fsnotify.Remove) {
				thread <- true
			} else if event.Has(fsnotify.Rename) {
				thread <- true
			}*/
		case err, ok := <-Watch.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
