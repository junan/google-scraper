package helpers_test

import (
	"testing"

	_ "google-scraper/initializers"
	. "google-scraper/tests/testing_helpers"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var _ = BeforeSuite(func() {
	web.TestBeegoInit(AppRootDir(1))
})
