package scraper

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beego/beego/v2/core/logs"
)

const googleSearchBaseUrl = "https://www.google.com/search"

type CrawlData struct {
	TopAdWordAdvertisersCount   int
	TopAdWordAdvertisersUrls    []string
	TotalAdWordAdvertisersCount int
	ResultsCount                int
	ResultsUrls                 []string
	TotalLinksCount             int
	Html                        string
}

func Crawl(searchString string) (data *CrawlData) {
	searchUrl := buildSearchUrl(searchString)
	response, err := getRequest(searchUrl)

	if err != nil {
		logs.Error(err)
	}

	htmlResponse := string(response)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlResponse))

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

func buildSearchUrl(searchString string) string {
	baseUrl, err := url.Parse(googleSearchBaseUrl)
	if err != nil {
		logs.Error("Parsing base user failed: ", err)
	}

	params := url.Values{}
	params.Add("q", searchString)
	params.Add("lr", "lang_en")
	params.Add("hl", "en")
	baseUrl.RawQuery = params.Encode()
	return baseUrl.String()
}
