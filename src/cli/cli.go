package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/BlazingFire007/blast/src/fileManager"
	"github.com/BlazingFire007/blast/src/logger"
	"github.com/urfave/cli/v2"
)

type FlagResults struct {
	Port int
	Dir  string
	Hot  bool
}

func (c *FlagResults) Init() {
	app := &cli.App{
		Name:  "blast",
		Usage: "put a folder on blast",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   3000,
				Usage:   "port to serve on",
			},
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Value:   ".",
				Usage:   "directory of static files to host",
			},
			&cli.BoolFlag{
				Name:    "refresh",
				Aliases: []string{"r"},
				Value:   false,
				Usage:   "disable hot refreshing",
			},
		},
		Action: func(cCtx *cli.Context) error {
			port := cCtx.Int("port")
			dir := cCtx.String("dir")
			port, dir = verify(port, dir)

			hot := cCtx.Bool("refresh")

			c.setData(port, dir, hot)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func (c *FlagResults) setData(port int, dir string, hot bool) {
	c.Port = port
	c.Dir = dir
	c.Hot = !hot
}

func verify(port int, dir string) (int, string) {
	if port < 1 || port > 65535 {
		logger.Error(fmt.Sprintf("Invalid port: %d", port))
	}
	exists := fileManager.CheckDir(dir)
	if !exists {
		logger.Error(fmt.Sprintf("Directory %s does not exist", dir))
	}
	dir = fileManager.FindDir(dir)
	return port, dir
}
