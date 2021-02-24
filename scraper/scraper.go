package scraper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beego/beego/v2/core/logs"
)

const googleSearchUrl = "https://www.google.com/search?q=%s&lr=lang_en"

type CrawlData struct {
	TopAdWordAdvertisersCount   int
	TopAdWordAdvertisersUrls    []string
	TotalAdWordAdvertisersCount int
	NonAdWordResultsCount       int
	NonAdWordResultsUrls        []string
	TotalLinksCount             int
	Html                        string
}

func Crawl(keyword string) (data *CrawlData) {
	url := generateTheQueryString(keyword)
	response, err := getRequest(url)
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
		TotalAdWordAdvertisersCount: 200,
		NonAdWordResultsCount:       300,
		NonAdWordResultsUrls:        []string{},
		TotalLinksCount:             len(getLinks(doc, "a")),
		Html:                        htmlResponse,
	}

	return data
}

func generateTheQueryString(keyword string) string {
	return fmt.Sprintf(googleSearchUrl, url.QueryEscape(keyword))
}

func getTopAdWordAdvertisersCount(doc *goquery.Document) int {
	return doc.Find("#tads > div").Length()
}

func getTopAdWordAdvertisersUrls(doc *goquery.Document) []string {
	return getLinks(doc, "#tads > div .Krnil")
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
