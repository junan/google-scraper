package crawler_test

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	_ "google-scraper/initializers"

	"github.com/beego/beego/v2/server/web"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrawler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crawler Suite")
}

var _ = BeforeSuite(func() {
	web.TestBeegoInit(JppRootDir(0))
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
