package models_test

import (
	_ "google-scraper/initializers"
	"google-scraper/models"
	. "google-scraper/tests/fabricators"
	. "google-scraper/tests/testing_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword", func() {
	Describe("#CreateKeyword", func() {
		Context("given the keyword with valid params", func() {
			It("returns the keyword ID", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := &models.Keyword{
					Name: "Buy Domain",
					User: &user,
				}
				keywordId, err := models.CreateKeyword(keyword)
				if err != nil {
					Fail("Storing keyword failed: " + err.Error())
				}

				Expect(keywordId).ToNot(BeNil())
			})

			It("returns an empty error", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := &models.Keyword{
					Name: "Buy Domain",
					User: &user,
				}
				_, err := models.CreateKeyword(keyword)

				Expect(err).To(BeNil())
			})
		})

		Context("given the keyword with INVALID params", func() {
			Context("given the user is nil", func() {
				It("returns an error", func() {
					keyword := &models.Keyword{
						Name: "Buy Domain",
						User: nil,
					}
					_, err := models.CreateKeyword(keyword)

					Expect(err.Error()).To(Equal("field `google-scraper/models.Keyword.User` cannot be NULL"))
				})
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
	})
})
