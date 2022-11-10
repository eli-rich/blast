package serve

import (
	"io/fs"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type Absolute string
type Relative string

var APP *fiber.App
var ROOT Absolute
var ROUTES map[Relative]Absolute

func initializeRouter(dir string, router *fiber.App) {
	APP = router
	ROOT = abs(Relative(dir))
	ROUTES = make(map[Relative]Absolute)
}

func initializePaths(dir string) {
	err := filepath.WalkDir(dir, visitor)
	panicOn(err)
}

func visitor(pathname string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if entry.IsDir() {
		return nil
	}
	filePath := abs(Relative(pathname))
	routePath := getRelFromRoot(filePath)
	if routePath == "/index.html" {
		routePath = "/"
	}
	ROUTES[routePath] = filePath
	return nil
}
