package jobs_test

import (
	_ "google-scraper/initializers"
	"google-scraper/models"
	. "google-scraper/tests"
	jobs "google-scraper/worker/jobs"

	work "github.com/gocraft/work"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crawling", func() {
	Describe("#PerformCrawling", func() {
		Context("given the params are valid", func() {
			It("returns no errors", func() {
				searchString := "Buy domain"
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword(searchString, false, &user)
				mockResponseFilePath := AppRootDir() + "/tests/fixtures/services/crawler/valid_get_response.html"
				MockCrawling(mockResponseFilePath)

				context := jobs.Context{}
				keywordJob := &work.Job{
					Args: work.Q{"keywordId": keyword.Id},
				}

				err := context.PerformCrawling(keywordJob)

				Expect(err).To(BeNil())
			})

			It("updates the keyword as completed", func() {
				searchString := "Buy domain"
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword(searchString, false, &user)
				mockResponseFilePath := AppRootDir() + "/tests/fixtures/services/crawler/valid_get_response.html"
				MockCrawling(mockResponseFilePath)

				context := jobs.Context{}
				keywordJob := &work.Job{
					Args: work.Q{"keywordId": keyword.Id},
				}

				err := context.PerformCrawling(keywordJob)
				if err != nil {
					Fail("Crawling failed: " + err.Error())
				}

				updatedKeyword, err := models.FindKeywordById(keyword.Id)
				if err != nil {
					Fail("Finding keyword failed: " + err.Error())
				}

				Expect(updatedKeyword.SearchCompleted).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})

		Context("given the params are INVALID", func() {
			It("returns an error", func() {
				context := jobs.Context{}
				keywordJob := &work.Job{
					Args: work.Q{"keywordId": -100},
				}

				err := context.PerformCrawling(keywordJob)

				Expect(err.Error()).To(Equal("<QuerySeter> no row found"))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords", "search_results")
	})
})
