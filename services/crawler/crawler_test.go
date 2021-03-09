package crawler_test

import (
	"encoding/json"

	. "google-scraper/services/crawler"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crawler", func() {
	Describe("#Crawl", func() {
		Context("given a valid search string", func() {
			It("returns the expected crawled data", func() {
				searchString := "Buy domain"
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword(searchString, &user)
				mockResponseFilePath := AppRootDir(0) + "/fixtures/services/crawler/valid_get_response.html"
				MockCrawling(mockResponseFilePath)

				data, err := Crawl(&keyword)
				if err != nil {
					Fail("Crawling failed: " + err.Error())
				}

				if err != nil {
					Fail("Crawling failed: " + err.Error())
				}

				var TopAdWordAdvertisersUrls []string
				err = json.Unmarshal([]byte(data.TopAdWordAdvertisersUrls), &TopAdWordAdvertisersUrls)
				if err != nil {
					Fail("Parsing TopAdWordAdvertisersUrls failed: " + err.Error())
				}

				var ResultsUrls []string
				err = json.Unmarshal([]byte(data.ResultsUrls), &ResultsUrls)
				if err != nil {
					Fail("Parsing ResultsUrls failed: " + err.Error())
				}

				Expect(data.TopAdWordAdvertisersCount).To(Equal(3))
				Expect(len(TopAdWordAdvertisersUrls)).To(Equal(3))
				Expect(data.TotalAdWordAdvertisersCount).To(Equal(3))
				Expect(data.ResultsCount).To(Equal(10))
				Expect(len(ResultsUrls)).To(Equal(10))
				Expect(data.TotalLinksCount).To(Equal(95))
				Expect(data.Html).NotTo(BeNil())
			})
		})

		Context("given an INVALID search string", func() {
			It("returns the expected crawled data", func() {
				searchString := ""
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword(searchString, &user)
				mockResponseFilePath := AppRootDir(0) + "/fixtures/services/crawler/invalid_get_response.html"
				MockCrawling(mockResponseFilePath)

				data, err := Crawl(&keyword)
				if err != nil {
					Fail("Crawling failed: " + err.Error())
				}

				if err != nil {
					Fail("Crawling failed: " + err.Error())
				}

				var TopAdWordAdvertisersUrls []string
				err = json.Unmarshal([]byte(data.TopAdWordAdvertisersUrls), &TopAdWordAdvertisersUrls)
				if err != nil {
					Fail("Parsing TopAdWordAdvertisersUrls failed: " + err.Error())
				}

				var ResultsUrls []string
				err = json.Unmarshal([]byte(data.ResultsUrls), &ResultsUrls)
				if err != nil {
					Fail("Parsing ResultsUrls failed: " + err.Error())
				}

				Expect(data.TopAdWordAdvertisersCount).To(Equal(0))
				Expect(len(TopAdWordAdvertisersUrls)).To(Equal(0))
				Expect(data.TotalAdWordAdvertisersCount).To(Equal(0))
				Expect(data.ResultsCount).To(Equal(0))
				Expect(len(ResultsUrls)).To(Equal(0))
				Expect(data.TotalLinksCount).To(Equal(17))
				Expect(data.Html).NotTo(BeNil())
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords", "search_results")
	})
})
