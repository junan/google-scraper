package testing_helpers

import (
	"net/http"
	"path"
	"path/filepath"
	"runtime"

	"github.com/onsi/ginkgo"
)

func AppRootDir(skip int) string {
	_, currentFile, _, _ := runtime.Caller(skip)
	currentFilePath := path.Join(path.Dir(currentFile))
	return filepath.Dir(currentFilePath)
}

func GetUrlPath(response *http.Response) string {
	url, err := response.Location()
	if err != nil {
		ginkgo.Fail("Getting current path failed: " + err.Error())
	}
	return url.Path
}
