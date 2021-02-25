package crawler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "google-scraper/tests/testing_helpers"

	. "google-scraper/services/crawler"
)

var _ = Describe("Crawler", func() {
	Describe("#Crawl", func() {
		BeforeEach(func() {
			RecordCassette("success_crawling", "Buy domain")
		})

		It("returns the crawled data", func() {
			data, err := Crawl("buy domain")
			if err != nil {
				Fail("Crawling failed: " + err.Error())
			}

			Expect(data.TopAdWordAdvertisersCount).To(Equal(4))
			Expect(len(data.TopAdWordAdvertisersUrls)).To(Equal(4))
			Expect(data.TotalAdWordAdvertisersCount).To(Equal(4))
			Expect(data.ResultsCount).To(Equal(10))
			Expect(len(data.ResultsUrls)).To(Equal(10))
			Expect(data.TotalLinksCount).To(Equal(94))
			Expect(data.Html).NotTo(BeNil())
		})

	})
})
