package enqueueing_test

import (
	"os"
	"path"
	"testing"

	_ "google-scraper/initializers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQueueing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queueing Suite")
}

var _ = BeforeSuite(func() {
	pwd, err := os.Getwd()
	if err != nil {
		logs.Error("Getting current directory failed", err)
	}

	web.TestBeegoInit(path.Join(pwd, "../.."))
})
