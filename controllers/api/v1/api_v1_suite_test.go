package apiv1_test

import (
	"os"
	"path"
	"testing"

	_ "google-scraper/routers"
	_ "google-scraper/initializers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAPIV1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API V1 Suite")
}

var _ = BeforeSuite(func() {
	pwd, err := os.Getwd()
	if err != nil {
		logs.Error("Getting current directory failed", err)
	}

	web.TestBeegoInit(path.Join(pwd, "../../.."))
})
