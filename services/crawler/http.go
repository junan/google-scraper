package crawler

import (
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/junan/fake-useragent"
)

func GetRequest(url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", randomUserAgent())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Closing response body
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Returns random Chrome or Firefox user agent.
// Browser version, system-information, platform also random.
func randomUserAgent() string {
	browserName := randomBrowser()
	switch browserName {
	case "chrome":
		return browser.Chrome()
	case "firefox":
		return browser.Firefox()
	default:
		return browser.Chrome()
	}
}

func randomBrowser() string {
	browsers := []string{
		"chrome",
		"firefox",
	}

	return browsers[rand.Intn(len(browsers))]
}
