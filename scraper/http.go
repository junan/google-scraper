package scraper

import (
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/junan/fake-useragent"
)

func getRequest(url string) ([]byte, error, string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logs.Error("Building new request failed: ", err)
	}

	agent := userAgent()
	req.Header.Set("User-Agent", agent)

	res, err := client.Do(req)
	if err != nil {
		return nil, err, agent
	}
	defer res.Body.Close()
	byte, err := ioutil.ReadAll(res.Body)

	return byte, err, agent
}

func userAgent() string {
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
