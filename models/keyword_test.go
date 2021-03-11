package models_test

import (
	_ "google-scraper/initializers"
	"google-scraper/models"
	. "google-scraper/tests"

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

			It("returns empty error", func() {
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

	Describe("#FindKeywordById", func() {
		Context("given the keyword already exist", func() {
			It("returns the keyword object", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				existingKeyword := FabricateKeyword("Buy domain", &user)
				keyword, err :=  models.FindKeywordById(existingKeyword.Id)
				if err != nil {
					Fail("Finding user failed: " + err.Error())
				}

				Expect(keyword.Id).To(Equal(existingKeyword.Id))
			})
		})

		Context("given the keyword does NOT exist", func() {
			It("returns an error", func() {
				_, err := models.FindKeywordById(-10)

				Expect(err.Error()).To(ContainSubstring("no row found"))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
	})
})
