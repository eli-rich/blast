package main

import (
	"github.com/BlazingFire007/blast/src/cli"
	"github.com/BlazingFire007/blast/src/serve"
)

func main() {
	c := cli.FlagResults{}
	c.Init()
	serve.Serve(c.Port, c.Dir)
}
