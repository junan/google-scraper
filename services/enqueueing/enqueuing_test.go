package enqueueing_test

import (
	"google-scraper/models"
	. "google-scraper/services/enqueueing"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Queueing", func() {
	Describe("#EnqueueKeywordJob", func() {
		Context("given a valid keyword object", func() {
			It("does NOT return error", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", &user)

				_, err := EnqueueKeywordJob(&keyword, 1)

				Expect(err).To(BeNil())
			})

			It("enqueues the job", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", &user)

				job, err := EnqueueKeywordJob(&keyword, 1)
				if err != nil {
					Fail("Adding to queue failed")
				}

				Expect(job.Name).To(Equal("crawling_job"))
				Expect(job.EnqueuedAt).NotTo(BeNil())
				Expect(job.EnqueuedAt).NotTo(BeNil())
			})
		})

		Context("given an empty keyword object", func() {
			It("returns an error", func() {
				job, err := EnqueueKeywordJob(&models.Keyword{}, 1)

				Expect(err.Error()).To(Equal("invalid keyword object"))
				Expect(job).To(BeNil())
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
		DeleteRedisJobs("google_scraper", "crawling_job")
	})
})
