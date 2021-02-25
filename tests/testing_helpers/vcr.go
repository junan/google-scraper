package testing_helpers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/onsi/ginkgo"
)

func RecordCassette(cassetteName string, searchString string) {
	searchUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en&hl=en", url.QueryEscape(searchString))
	recorder, err := recorder.New(fmt.Sprintf("%s/tests/fixtures/vcr/%s", AppRootDir(), cassetteName))
	if err != nil {
		Fail(err.Error())
	}
	defer recorder.Stop()

	client := &http.Client{Transport: recorder}
	_, err = client.Get(searchUrl)
	if err != nil {
		Fail(err.Error())
	}
}
