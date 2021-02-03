package controllers_test

import (
	"testing"

	_ "google-scraper/initializers"
	_ "google-scraper/routers"
	. "google-scraper/tests/testing_helpers"

	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var _ = BeforeSuite(func() {
	web.TestBeegoInit(AppRootDir())
})
