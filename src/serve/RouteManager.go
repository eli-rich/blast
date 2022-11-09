package serve

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/BlazingFire007/blast/src/refresh"
	"github.com/gin-gonic/gin"
)

type Route struct {
	RelativePath string
	PathName     string
}

var ROUTER *gin.Engine
var ROUTES []Route

var pathnames []string
var ROOT string

func Initialize(router *gin.Engine, dir string) {
	ROUTER = router
	ROOT = dir
}

func AddRoutes(dir string) {
	filepath.WalkDir(dir, visit)
	for _, pathname := range pathnames {
		createRoute(pathname)
	}
}

func visit(pathname string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		pathnames = append(pathnames, pathname)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}

func createRoute(pathname string) {
	if filepath.Base(pathname) == "index.html" {
		refresh.InjectScript(ROUTER, pathname, HOT)
	} else {
		var relativePath string
		if ROOT == "." {
			relativePath = pathname
		} else {
			relativePath = strings.Replace(pathname, ROOT, "", 1)
		}
		route := Route{RelativePath: relativePath, PathName: pathname}
		ROUTES = append(ROUTES, route)
		ROUTER.StaticFile(route.RelativePath, route.PathName)
	}
}
