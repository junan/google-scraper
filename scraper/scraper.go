package scraper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beego/beego/v2/core/logs"
)

const googleSearchUrl = "https://www.google.com/search?q=rails&lr=lang_en&hl=en"

type CrawlData struct {
	TopAdWordAdvertisersCount   int
	TopAdWordAdvertisersUrls    []string
	TotalAdWordAdvertisersCount int
	ResultsCount                int
	ResultsUrls                 []string
	TotalLinksCount             int
	Html                        string
}

func Crawl(keyword string) (data *CrawlData) {
	response, err := getRequest(googleSearchUrl)
	if err != nil {
		logs.Error(err)
	}

	htmlResponse := string(response)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlResponse))

	result := len(getLinks(doc, "a"))

	fmt.Printf("", result)

	if err != nil {
		logs.Error(err)
	}

	data = &CrawlData{
		TopAdWordAdvertisersCount:   getTopAdWordAdvertisersCount(doc),
		TopAdWordAdvertisersUrls:    getTopAdWordAdvertisersUrls(doc),
		TotalAdWordAdvertisersCount: GetTotalAdWordAdvertisersCount(doc),
		ResultsCount:                getResultsCount(doc),
		ResultsUrls:                 getResultsUrls(doc),
		TotalLinksCount:             len(getLinks(doc, "a")),
		Html:                        htmlResponse,
	}

	return data
}

func generateTheQueryString(keyword string) string {
	return fmt.Sprintf(googleSearchUrl, url.QueryEscape(keyword))
}

func getTopAdWordAdvertisersCount(doc *goquery.Document) int {
	return doc.Find("#tads .uEierd").Length()
}

func GetTotalAdWordAdvertisersCount(doc *goquery.Document) int {
	return len(getLinks(doc, ".Krnil"))
}

func getTopAdWordAdvertisersUrls(doc *goquery.Document) []string {
	return getLinks(doc, "#tads .Krnil")
}

func getResultsCount(doc *goquery.Document) int {
	return doc.Find("#rso .yuRUbf").Length()
}

func getResultsUrls(doc *goquery.Document) []string {
	return getLinks(doc, "#rso .yuRUbf > a")
}

func getLinks(doc *goquery.Document, selector string) []string {
	var links []string

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
		}
	})

	return links
}
