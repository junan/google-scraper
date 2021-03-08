package crawler_test

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	_ "google-scraper/initializers"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrawler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crawler Suite")
}

var _ = BeforeSuite(func() {
	pwd, err := os.Getwd()
	if err != nil {
		logs.Error("Getting current directory failed", err)
	}

	web.TestBeegoInit(path.Join(pwd, "../.."))
	// block all HTTP requests
	httpmock.Activate()
})

var _ = BeforeEach(func() {
	// remove any mocks
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})

func JppRootDir(skip int) string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentFilePath := path.Join(path.Dir(currentFile), "..")
	result := filepath.Dir(currentFilePath)
	return result
}
