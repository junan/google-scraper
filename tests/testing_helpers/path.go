package test_helpers

import (
	"path"
	"path/filepath"
	"runtime"
)

func AppRootDir() string {
	_, currentFile, _, _ := runtime.Caller(1)
	currentFilePath := path.Join(path.Dir(currentFile))
	return filepath.Dir(currentFilePath)
}
