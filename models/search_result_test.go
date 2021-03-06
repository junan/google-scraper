package models_test

import (
	_ "google-scraper/initializers"
	"google-scraper/models"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchResult", func() {
	Describe("#CreateSearchResult", func() {
		Context("given the SearchResult with valid params", func() {
			It("returns the search result ID", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", false, &user)
				searchResult := &models.SearchResult{
					TopAdWordAdvertisersCount:   2,
					TopAdWordAdvertisersUrls:    `["http://example1.com", "http://example2.com"]`,
					TotalAdWordAdvertisersCount: 3,
					ResultsCount:                2,
					ResultsUrls:                 `["http://example1.com", "http://example2.com"]`,
					TotalLinksCount:             20,
					Html:                        "html-response",
					Keyword:                     &keyword,
				}
				searchResultId, err := models.CreateSearchResult(searchResult)
				if err != nil {
					Fail("Storing search result failed: " + err.Error())
				}

				Expect(searchResultId).ToNot(BeNil())
			})

			It("returns empty error", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", false, &user)
				searchResult := &models.SearchResult{
					TopAdWordAdvertisersCount:   2,
					TopAdWordAdvertisersUrls:    `["http://example1.com", "http://example2.com"]`,
					TotalAdWordAdvertisersCount: 3,
					ResultsCount:                2,
					ResultsUrls:                 `["http://example1.com", "http://example2.com"]`,
					TotalLinksCount:             20,
					Html:                        "html-response",
					Keyword:                     &keyword,
				}
				_, err := models.CreateSearchResult(searchResult)

				Expect(err).To(BeNil())
			})
		})

		Context("given the SearchResult with INVALID params", func() {
			Context("given the keyword is nil", func() {
				It("returns an error", func() {
					user := FabricateUser("John", "john@example.com", "secret")
					keyword := FabricateKeyword("Buy domain", false, &user)
					searchResult := &models.SearchResult{
						TopAdWordAdvertisersCount:   2,
						TopAdWordAdvertisersUrls:    `["http://example1.com", "http://example2.com"]`,
						TotalAdWordAdvertisersCount: 3,
						ResultsCount:                2,
						ResultsUrls:                 "[http://example1.com, http://example2.com]",
						TotalLinksCount:             20,
						Html:                        "html-response",
						Keyword:                     &keyword,
					}
					_, err := models.CreateSearchResult(searchResult)

					Expect(err.Error()).To(Equal("pq: invalid input syntax for type json"))
				})
			})
		})
	})

	Describe("#FindSearchResultByKeywordId", func() {
		Context("given the keyword already exists", func() {
			It("returns the keyword's SearchResult object", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", false, &user)
				existingSearchResult := FabricateSearchResult(&keyword)
				searchResult, err := models.FindSearchResultByKeywordId(keyword.Id)
				if err != nil {
					Fail("Finding search result failed: " + err.Error())
				}

				Expect(searchResult.Id).To(BeNumerically("==", existingSearchResult.Id))
			})
		})

		Context("given the keyword does NOT exist", func() {
			It("returns an error", func() {
				_, err := models.FindSearchResultByKeywordId(1000)

				Expect(err.Error()).To(Equal("<QuerySeter> no row found"))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords", "search_results")
	})
})
