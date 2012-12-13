package goit

import (
	"path/filepath"
)

func cleanupPath(path string) string {
	path, _ = filepath.EvalSymlinks(path)
	return path
}
