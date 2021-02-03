package controllers_test

import (
	"github.com/beego/beego/v2/server/web"
	_ "google-scraper/initializers"
	_ "google-scraper/routers"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

func appRootDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentFilePath := path.Join(path.Dir(currentFile))
	return filepath.Dir(currentFilePath)
}

var _ = BeforeSuite(func() {
	web.TestBeegoInit(appRootDir())
})
