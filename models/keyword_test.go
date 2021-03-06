package models_test

import (
	"fmt"

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

	Describe("#UpdateKeyword", func() {
		It("updates keyword status as completed", func() {
			user := FabricateUser("John", "john@example.com", "secret")
			keyword := FabricateKeyword("Buy domain", false, &user)
			keyword.SearchCompleted = true
			updatedKeyword, err := models.UpdateKeyword(&keyword)

			if err != nil {
				Fail("Updating keyword failed: " + err.Error())
			}

			Expect(updatedKeyword.SearchCompleted).To(BeTrue())
		})
	})

	Describe("#FindKeywordById", func() {
		Context("given the keyword already exist", func() {
			It("returns the keyword object", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				existingKeyword := FabricateKeyword("Buy domain", false, &user)
				keyword, err := models.FindKeywordById(existingKeyword.Id)
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

	Describe("#GetGetQuerySeterKeywords", func() {
		Context("given params with keyword filtering", func() {
			Context("given the filtered keyword exists in the database", func() {
				It("returns matched user keywords", func() {
					var keywordIds []int64
					userKeywords := []*models.Keyword{}
					keyword := "Buy domain"
					user := FabricateUser("John", "john@example.com", "secret")
					buyDomainKeyword := FabricateKeyword(keyword, false, &user)
					FabricateKeyword("Purchase domain", false, &user)
					FabricateKeyword("Buy bike", false, &user)
					user2 := FabricateUser("David", "david@example.com", "secret")
					FabricateKeyword("Buy car", false, &user2)

					_, err := models.GetQuerySeterKeywords(&user, keyword).All(&userKeywords)
					if err != nil {
						Fail("Getting keyword failed: " + err.Error())
					}

					for _, v := range userKeywords {
						keywordIds = append(keywordIds, v.Id)
					}

					expectedKeywordIds := []int64{buyDomainKeyword.Id}

					Expect(keywordIds).To(Equal(expectedKeywordIds))
				})
			})

			Context("given the filtered keyword does NOT exist in the database", func() {
				It("returns empty keywords", func() {
					nonExistedKeyword := "Non existed keyword"
					user := FabricateUser("John", "john@example.com", "secret")
					FabricateKeyword("Buy domain", false, &user)
					FabricateKeyword("Purchase domain", false, &user)
					FabricateKeyword("Buy bike", false, &user)
					user2 := FabricateUser("David", "david@example.com", "secret")
					FabricateKeyword("Buy car", false, &user2)

					keywords := models.GetQuerySeterKeywords(&user, nonExistedKeyword)

					Expect(keywords.Count()).To(BeNumerically("==", 0))
				})
			})
		})

		Context("given params without keyword filtering", func() {
			It("returns user keywords", func() {
				var expectedKeyword models.Keyword
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", false, &user)

				keywords := models.GetQuerySeterKeywords(&user, "")
				err := keywords.One(&expectedKeyword)
				if err != nil {
					Fail("Getting keyword failed: " + err.Error())
				}

				Expect(keywords.Count()).To(BeNumerically("==", 1))
				Expect(keyword.Id).To(BeNumerically("==", expectedKeyword.Id))
			})
		})
	})

	Describe("#GetPaginatedKeywords", func() {
		Context("given the sizePerPage is 10", func() {
			It("returns 10 keywords", func() {
				user := FabricateUser("John", "john@example.com", "secret")

				// Creating 11 keywords record
				for i := 1; i <= 11; i++ {
					FabricateKeyword(fmt.Sprintf("Buy domain %d", i), false, &user)
				}

				keywords := models.GetQuerySeterKeywords(&user, "")

				paginatedKeywords, err := models.GetPaginatedKeywords(keywords, 0, 10)
				if err != nil {
					Fail("Getting paginated keywords failed: " + err.Error())
				}

				Expect(len(paginatedKeywords)).To(BeNumerically("==", 10))
			})
		})

		Context("given the sizePerPage is 5", func() {
			It("returns 5 keywords", func() {
				user := FabricateUser("John", "john@example.com", "secret")

				// Creating 6 keywords record
				for i := 1; i <= 6; i++ {
					FabricateKeyword(fmt.Sprintf("Buy domain %d", i), false, &user)
				}

				keywords := models.GetQuerySeterKeywords(&user, "")

				paginatedKeywords, err := models.GetPaginatedKeywords(keywords, 0, 5)
				if err != nil {
					Fail("Getting paginated keywords failed: " + err.Error())
				}

				Expect(len(paginatedKeywords)).To(BeNumerically("==", 5))
			})
		})
	})

	Describe("#FindKeywordBy", func() {
		Context("given the keyword belongs to the user", func() {
			It("returns the keyword", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				keyword := FabricateKeyword("Buy domain", false, &user)

				result, err := models.FindKeywordBy(keyword.Id, &user)
				if err != nil {
					Fail("Finding keyword failed: " + err.Error())
				}

				Expect(result.Id).To(Equal(keyword.Id))
			})
		})

		Context("given the keyword does NOT belong to the user", func() {
			It("returns an error", func() {
				user1 := FabricateUser("John", "john@example.com", "secret")
				FabricateKeyword("Buy domain", false, &user1)
				user2 := FabricateUser("Mike", "mike@example.com", "secret")
				keyword2 := FabricateKeyword("Buy bike", false, &user2)

				_, err := models.FindKeywordBy(keyword2.Id, &user1)

				Expect(err.Error()).To(Equal("Keyword not found."))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
	})
})
