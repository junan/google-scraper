package jobs_test

import (
	"os"
	"path"
	"testing"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJobs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jobs Suite")
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
