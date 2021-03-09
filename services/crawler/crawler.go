package crawler

import (
	"strings"

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

type CrawlData struct {
	TopAdWordAdvertisersCount   int
	TopAdWordAdvertisersUrls    []string
	TotalAdWordAdvertisersCount int
	ResultsCount                int
	ResultsUrls                 []string
	TotalLinksCount             int
	Html                        string
}

// Call this function with your search key, it will return the necessary crawled data
// Ex: Crawl("Buy laptop")
func Crawl(searchString string) (data *CrawlData, err error) {
	searchUrl, err := BuildSearchUrl(searchString)
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

	data = parseCrawledData()

	return data, nil
}

func getTopAdWordAdvertisersCount() int {
	return doc.Find(selectorMapping["topAdWordAdvertisersCount"]).Length()
}

func getTopAdWordAdvertisersUrls() []string {
	return getLinks(selectorMapping["topAdWordAdvertisersUrls"])
}

func GetTotalAdWordAdvertisersCount() int {
	return len(getLinks(selectorMapping["totalAdWordAdvertisersCount"]))
}

func getResultsCount() int {
	return doc.Find(selectorMapping["resultsCount"]).Length()
}

func getResultsUrls() []string {
	return getLinks(selectorMapping["resultsUrls"])
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

func parseCrawledData() *CrawlData {
	return &CrawlData{
		TopAdWordAdvertisersCount:   getTopAdWordAdvertisersCount(),
		TopAdWordAdvertisersUrls:    getTopAdWordAdvertisersUrls(),
		TotalAdWordAdvertisersCount: GetTotalAdWordAdvertisersCount(),
		ResultsCount:                getResultsCount(),
		ResultsUrls:                 getResultsUrls(),
		TotalLinksCount:             getTotalLinks(),
		Html:                        htmlResponse,
	}
}
