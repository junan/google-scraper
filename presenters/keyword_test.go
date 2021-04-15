package presenters_test

import (
	"encoding/json"

	_ "google-scraper/initializers"
	"google-scraper/presenters"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword", func() {
	Describe("#InitializeKeywordPresenter", func() {
		Context("given the valid keyword", func() {
			It("returns the presenter object", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", false, &user)
				searchResult := FabricateSearchResult(&keyword)

				presenter, err := presenters.InitializeKeywordPresenter(&keyword)
				if err != nil {
					Fail("Initializing presenter failed: " + err.Error())
				}

				var topAdWordAdvertisersUrls []string
				err = json.Unmarshal([]byte(searchResult.TopAdWordAdvertisersUrls), &topAdWordAdvertisersUrls)
				if err != nil {
					Fail("Unmarshal failed: " + err.Error())
				}

				var resultsUrls []string
				err = json.Unmarshal([]byte(searchResult.ResultsUrls), &resultsUrls)
				if err != nil {
					Fail("Unmarshal failed: " + err.Error())
				}

				Expect(presenter).NotTo(BeNil())
				Expect(presenter.Keyword).To(Equal(keyword.Name))
				Expect(presenter.CreatedAt).NotTo(BeEmpty())
				Expect(presenter.TopAdWordAdvertisersCount).To(Equal(searchResult.TopAdWordAdvertisersCount))
				Expect(presenter.TotalAdWordAdvertisersCount).To(Equal(searchResult.TotalAdWordAdvertisersCount))
				Expect(presenter.ResultsCount).To(Equal(searchResult.ResultsCount))
				Expect(presenter.TotalLinksCount).To(Equal(searchResult.TotalLinksCount))
				Expect(presenter.Html).To(Equal(searchResult.Html))
				Expect(presenter.TopAdWordAdvertisersUrls).To(Equal(topAdWordAdvertisersUrls))
				Expect(presenter.ResultsUrls).To(Equal(resultsUrls))
			})
		})

		Context("given the INVALID keyword", func() {
			It("returns an error", func() {
				_, err := presenters.InitializeKeywordPresenter(nil)

				Expect(err.Error()).To(Equal("keyword object can't be nil"))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
	})
})
