package main

import (
	"github.com/BlazingFire007/blast/src/cli"
	"github.com/BlazingFire007/blast/src/graceful"
	"github.com/BlazingFire007/blast/src/refresh"
	"github.com/BlazingFire007/blast/src/serve"
	"github.com/BlazingFire007/blast/src/watcher"
)

func main() {
	c := cli.FlagResults{}
	c.Init()
	graceful.Graceful(c.Dir, c.Hot)
	refresh.CreateRefresher(c.Hot, c.Dir)
	watcher.Init(c.Hot)
	serve.Serve(c.Port, c.Dir, c.Hot)
}
