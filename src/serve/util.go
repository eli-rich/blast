package serve

import (
	"log"
	"path/filepath"
)

func getRelFromRoot(pathname Absolute) Relative {
	return Relative(pathname[len(ROOT):])
}

func abs(pathname Relative) Absolute {
	abs, err := filepath.Abs(string(pathname))
	panicOn(err)
	return Absolute(abs)
}

func panicOn(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
