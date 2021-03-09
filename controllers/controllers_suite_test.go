package controllers_test

import (
	"testing"

	_ "google-scraper/initializers"
	_ "google-scraper/routers"
	. "google-scraper/tests"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jarcoal/httpmock"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	web.TestBeegoInit(AppRootDir(1))
	httpmock.Activate()
})

var _ = BeforeEach(func() {
	// remove any mocks
	httpmock.Reset()
})

var _ = AfterSuite(func() {
	httpmock.DeactivateAndReset()
})
