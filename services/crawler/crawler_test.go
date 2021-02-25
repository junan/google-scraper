package crawler_test

import (
	"io/ioutil"
	"fmt"

	. "google-scraper/services/crawler"
	. "google-scraper/tests/testing_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jarcoal/httpmock"
)

var _ = Describe("Crawler", func() {
	Describe("#Crawl", func() {
		It("returns the crawled data", func() {
			searchUrl := BuildSearchUrl("Buy domain")
			htmlPath := fmt.Sprintf("%s/fixtures/buy_domain.html", AppRootDir(0))

			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			content, err := ioutil.ReadFile(htmlPath)
			if err != nil {
				Fail("Reading file failed: " + err.Error())
			}

			httpmock.RegisterResponder("GET", searchUrl,
				httpmock.NewStringResponder(200, string(content)))

			data, err := Crawl("Buy domain")
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
})
