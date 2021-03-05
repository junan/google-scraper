package crawler_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Crawler", func() {
	Describe("#Crawl", func() {
		Context("given a valid search string", func() {
			//It("returns the expected crawled data", func() {
			//	searchString := "Buy domain"
			//	htmlPath := fmt.Sprintf("%s/fixtures/services/crawler/valid_get_request.html", AppRootDir(0))
			//	user := FabricateUser("John", "john@example.com", "secret")
			//	keyword := FabricateKeyword(searchString, &user)
			//	MockCrawling(keyword.Name, htmlPath)
			//
			//	data, err := Crawl(keyword)
			//	if err != nil {
			//		Fail("Crawling failed: " + err.Error())
			//	}
			//
			//
			//	if err != nil {
			//		Fail("Crawling failed: " + err.Error())
			//	}
			//
			//	Expect(data.TopAdWordAdvertisersCount).To(Equal(3))
			//	Expect(len(data.TopAdWordAdvertisersUrls)).To(Equal(3))
			//	Expect(data.TotalAdWordAdvertisersCount).To(Equal(3))
			//	Expect(data.ResultsCount).To(Equal(10))
			//	Expect(len(data.ResultsUrls)).To(Equal(10))
			//	Expect(data.TotalLinksCount).To(Equal(95))
			//	Expect(data.Html).NotTo(BeNil())
			//})
		})

		Context("given an INVALID search string", func() {
			//It("returns the expected crawled data", func() {
			//	searchString := ""
			//	htmlPath := fmt.Sprintf("%s/fixtures/services/crawler/invalid_get_request.html", AppRootDir(0))
			//	user := FabricateUser("John", "john@example.com", "secret")
			//	keyword := FabricateKeyword(searchString, &user)
			//	MockCrawling(keyword.Name, htmlPath)
			//
			//	data, err := Crawl(keyword)
			//	if err != nil {
			//		Fail("Crawling failed: " + err.Error())
			//	}
			//
			//	Expect(data.TopAdWordAdvertisersCount).To(Equal(0))
			//	Expect(len(data.TopAdWordAdvertisersUrls)).To(Equal(0))
			//	Expect(data.TotalAdWordAdvertisersCount).To(Equal(0))
			//	Expect(data.ResultsCount).To(Equal(0))
			//	Expect(len(data.ResultsUrls)).To(Equal(0))
			//	Expect(data.TotalLinksCount).To(Equal(17))
			//	Expect(data.Html).NotTo(BeNil())
			//})
		})
	})
})
