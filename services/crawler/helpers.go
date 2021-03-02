package crawler

import (
	"net/url"
)

func BuildSearchUrl(searchString string, rootUrl string) (string, error) {
	baseUrl, err := url.Parse(rootUrl)
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("q", searchString)
	params.Add("lr", "lang_en")
	params.Add("hl", "en")
	baseUrl.RawQuery = params.Encode()

	return baseUrl.String(), nil
}
