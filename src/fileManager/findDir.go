package fileManager

import (
	"os"
	"path"
)

// Check if directory exists.
func CheckDir(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

// Find index.html in directory. If not found, look in ./public.
func FindIndex(dir string) string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.Name() == "index.html" {
			return path.Join(dir, entry.Name())
		}
	}
	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() == "public" {
				return FindIndex(path.Join(dir, entry.Name()))
			} else if entry.Name() == "static" {
				return FindIndex(path.Join(dir, entry.Name()))
			} else if entry.Name() == "dist" {
				return FindIndex(path.Join(dir, entry.Name()))
			}
		}
	}
	return ""
}

// return base dir of index.html
func FindDir(dir string) string {
	index := FindIndex(dir)
	return path.Dir(index)
}
