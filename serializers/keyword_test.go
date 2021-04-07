package serializers_test

import (
	. "google-scraper/presenters"
	"google-scraper/serializers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordSerializer", func() {
	Describe("#GetKeywordResponse", func() {
		It("returns correct data", func() {
			searchResult := KeywordSearchResult{
				KeywordId:                   1,
				Keyword:                     "Buy domain",
				CreatedAt:                   "March 24, 2021",
				TopAdWordAdvertisersCount:   1,
				TotalAdWordAdvertisersCount: 2,
				TotalLinksCount:             20,
				ResultsCount:                10,
				Html:                        "html",
				TopAdWordAdvertisersUrls:    []string{"example.com"},
				ResultsUrls:                 []string{"example2.com", "example3.com"},
			}

			response := serializers.GetKeywordResponse(searchResult)

			Expect(response.Id).To(Equal(int64(1)))
			Expect(response.Name).To(Equal("Buy domain"))
			Expect(response.CreatedAt).To(Equal("March 24, 2021"))
			Expect(response.TopAdWordAdvertisersCount).To(Equal(1))
			Expect(response.TotalAdWordAdvertisersCount).To(Equal(2))
			Expect(response.TotalLinksCount).To(Equal(20))
			Expect(response.ResultsCount).To(Equal(10))
			Expect(response.Html).To(Equal("html"))
			Expect(response.TopAdWordAdvertisersUrls).To(Equal([]string{"example.com"}))
			Expect(response.ResultsUrls).To(Equal([]string{"example2.com", "example3.com"}))
		})
	})
})
