package crawler

import (
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/junan/fake-useragent"
)

func getRequest(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.Error("Building new request failed: ", err)
	}

	req.Header.Set("User-Agent", randomUserAgent())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Closing response body
	defer res.Body.Close()
	byte, err := ioutil.ReadAll(res.Body)

	return byte, err
}

// Return random Chrome or Firefox user agent
// Browser version, system-information, platform also random
func randomUserAgent() string {
	rangeLower := 0
	rangeUpper := 1
	randNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)

	switch randNum {
	case 0:
		return browser.Chrome()
	case 1:
		return browser.Firefox()
	default:
		return browser.Chrome()
	}
}
