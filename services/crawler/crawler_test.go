package crawler_test

import (
	"fmt"

	. "google-scraper/services/crawler"
	. "google-scraper/tests/testing_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crawler", func() {
	Describe("#Crawl", func() {
		Context("given a valid search string", func() {
			It("returns the expected crawled data", func() {
				searchString := "Buy domain"
				htmlPath := fmt.Sprintf("%s/fixtures/services/crawler/valid_get_request.html", AppRootDir(0))
				MockCrawling(searchString, htmlPath)

				data, err := Crawl(searchString)
				if err != nil {
					Fail("Crawling failed: " + err.Error())
				}

				Expect(data.TopAdWordAdvertisersCount).To(Equal(3))
				Expect(len(data.TopAdWordAdvertisersUrls)).To(Equal(3))
				Expect(data.TotalAdWordAdvertisersCount).To(Equal(3))
				Expect(data.ResultsCount).To(Equal(10))
				Expect(len(data.ResultsUrls)).To(Equal(10))
				Expect(data.TotalLinksCount).To(Equal(95))
				Expect(data.Html).NotTo(BeNil())
			})
		})

		Context("given an INVALID search string", func() {
			It("returns the expected crawled data", func() {
				searchString := ""
				htmlPath := fmt.Sprintf("%s/fixtures/services/crawler/invalid_get_request.html", AppRootDir(0))
				MockCrawling(searchString, htmlPath)

				data, err := Crawl(searchString)
				if err != nil {
					Fail("Crawling failed: " + err.Error())
				}

				Expect(data.TopAdWordAdvertisersCount).To(Equal(0))
				Expect(len(data.TopAdWordAdvertisersUrls)).To(Equal(0))
				Expect(data.TotalAdWordAdvertisersCount).To(Equal(0))
				Expect(data.ResultsCount).To(Equal(0))
				Expect(len(data.ResultsUrls)).To(Equal(0))
				Expect(data.TotalLinksCount).To(Equal(17))
				Expect(data.Html).NotTo(BeNil())
			})
		})
	})
})
