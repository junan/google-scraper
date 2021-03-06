package crawler

import (
	"encoding/json"
	"strings"

	"google-scraper/models"

	"github.com/PuerkitoBio/goquery"
)

const googleSearchBaseUrl = "https://www.google.com/search"

var doc *goquery.Document
var htmlResponse string
var selectorMapping = map[string]string{
	"topAdWordAdvertisersCount":   "#tads .uEierd",
	"topAdWordAdvertisersUrls":    "#tads .Krnil",
	"totalAdWordAdvertisersCount": ".Krnil",
	"resultsCount":                "#rso .yuRUbf",
	"resultsUrls":                 "#rso .yuRUbf > a",
	"totalLinksCount":             "a",
}

// Call this function with by passi g the keyword object, it will return the necessary crawled data
// Ex: Crawl(keywordObject)
func Crawl(keyword *models.Keyword) (searchResult *models.SearchResult, err error) {
	searchUrl, err := BuildSearchUrl(keyword.Name)
	if err != nil {
		return nil, err
	}

	response, err := GetRequest(searchUrl)
	if err != nil {
		return nil, err
	}

	htmlResponse = string(response)
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(htmlResponse))
	if err != nil {
		return nil, err
	}

	searchResult = &models.SearchResult{
		TopAdWordAdvertisersCount:   getTopAdWordAdvertisersCount(),
		TopAdWordAdvertisersUrls:    getTopAdWordAdvertisersUrls(),
		TotalAdWordAdvertisersCount: getTotalAdWordAdvertisersCount(),
		ResultsCount:                getResultsCount(),
		ResultsUrls:                 getResultsUrls(),
		TotalLinksCount:             getTotalLinks(),
		Html:                        htmlResponse,
		Keyword:                     keyword,
	}

	_, err = models.CreateSearchResult(searchResult)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

func getTopAdWordAdvertisersCount() int {
	return doc.Find(selectorMapping["topAdWordAdvertisersCount"]).Length()
}

func getTopAdWordAdvertisersUrls() string {
	return parsedUrls("topAdWordAdvertisersUrls")
}

func getTotalAdWordAdvertisersCount() int {
	return len(getLinks(selectorMapping["totalAdWordAdvertisersCount"]))
}

func getResultsCount() int {
	return doc.Find(selectorMapping["resultsCount"]).Length()
}

func getResultsUrls() string {
	return parsedUrls("resultsUrls")
}

func getTotalLinks() int {
	return len(getLinks(selectorMapping["totalLinksCount"]))
}

func getLinks(selector string) []string {
	var links []string

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
		}
	})

	return links
}

func parsedUrls(selector string) string {
	links := getLinks(selectorMapping[selector])
	urls, err := json.Marshal(links)
	if err != nil {
		return ""
	}

	return string(urls)
}
