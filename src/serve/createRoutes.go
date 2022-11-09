package serve

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/BlazingFire007/blast/src/refresh"
	"github.com/gin-gonic/gin"
)

var pathnames []string
var root string

func createRoutes(router *gin.Engine, dir string) {
	root = dir
	filepath.WalkDir(dir, visit)
	for _, pathname := range pathnames {
		if filepath.Base(pathname) == "index.html" {
			refresh.InjectScript(router, pathname, HOT)
		} else {
			var relativePath string
			if root == "." {
				relativePath = pathname
			} else {
				relativePath = strings.Replace(pathname, root, "", 1)
			}
			router.StaticFile(relativePath, pathname)
		}
	}
}

func visit(pathname string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		pathnames = append(pathnames, pathname)
	}
	return nil
}
