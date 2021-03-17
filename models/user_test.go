package models_test

import (
	"fmt"

	_ "google-scraper/initializers"
	"google-scraper/models"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("#Keywords", func() {
		It("returns user keywords", func() {
			user := FabricateUser("John", "john@example.com", "secret")
			FabricateKeyword("Buy domain", false, &user)
			FabricateKeyword("Buy bike", false, &user)
			user2 := FabricateUser("David", "david@example.com", "secret")
			FabricateKeyword("Buy car", false, &user2)

			Expect(user.Keywords().Count()).To(BeNumerically("==", 2))
		})
	})

	Describe("#PaginatedKeywords", func() {
		Context("given the sizePerPage is 10", func() {
			It("returns 10 keywords", func() {
				user := FabricateUser("John", "john@example.com", "secret")

				// Creating 11 keywords record
				for i := 1; i <= 11; i++ {
					FabricateKeyword(fmt.Sprintf("Buy domain %d", i), false, &user)
				}

				paginatedKeywords, err := user.PaginatedKeywords(user.Keywords(), 0, 10)
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

				paginatedKeywords, err := user.PaginatedKeywords(user.Keywords(), 0, 5)
				if err != nil {
					Fail("Getting paginated keywords failed: " + err.Error())
				}

				Expect(len(paginatedKeywords)).To(BeNumerically("==", 5))
			})
		})
	})

	Describe("#CreateUser", func() {
		Context("given the user with valid params", func() {
			It("returns the user ID", func() {
				user := &models.User{
					Name:           "John",
					Email:          "john@example.com",
					HashedPassword: "secret",
				}
				userId, err := models.CreateUser(user)
				if err != nil {
					Fail("Adding user failed: " + err.Error())
				}

				Expect(userId).ToNot(BeNil())
			})

			It("returns an empty error", func() {
				user := &models.User{
					Name:           "John",
					Email:          "john@example.com",
					HashedPassword: "secret",
				}
				_, err := models.CreateUser(user)

				Expect(err).To(BeNil())
			})
		})

		Context("given the user with INVALID params", func() {
			Context("given the email that is already exist", func() {
				It("returns an error", func() {
					email := "john@example.com"
					FabricateUser("John", email, "secret")

					user := &models.User{
						Name:           "John",
						Email:          email,
						HashedPassword: "secret",
					}
					_, err := models.CreateUser(user)

					Expect(err.Error()).To(Equal(`pq: duplicate key value violates unique constraint "users_email_key"`))
				})
			})
		})
	})

	Describe("#FindUserById", func() {
		Context("given the user already exist", func() {
			It("returns the user object", func() {
				existingUser := FabricateUser("John", "john@example.co", "secret")
				user, err := models.FindUserById(existingUser.Id)
				if err != nil {
					Fail("Finding user failed: " + err.Error())
				}

				Expect(user.Id).To(Equal(existingUser.Id))
			})
		})

		Context("given the user does NOT exist", func() {
			It("returns an error", func() {
				_, err := models.FindUserById(-10)

				Expect(err.Error()).To(ContainSubstring("no row found"))
			})
		})
	})

	Describe("#FindUserByEmail", func() {
		Context("given the user already exist", func() {
			It("returns the user", func() {
				email := "john@example.com"
				existingUser := FabricateUser("John", email, "secret")

				user, err := models.FindUserByEmail(email)
				if err != nil {
					Fail("Finding user failed: " + err.Error())
				}

				Expect(user.Email).To(Equal(existingUser.Email))
			})
		})

		Context("given the user does NOT exist", func() {
			It("returns an error", func() {
				_, err := models.FindUserByEmail("non_existing_email@example.com")

				Expect(err.Error()).To(ContainSubstring("no row found"))
			})
		})
	})

	Describe("#IsExistingUser", func() {
		Context("given the user already exist", func() {
			It("returns true", func() {
				email := "john@example.com"
				FabricateUser("John", email, "secret")
				user := &models.User{
					Email: email,
				}

				Expect(user.IsExistingUser()).To(BeTrue())
			})
		})

		Context("given the user does NOT exist", func() {
			It("returns false", func() {
				email := "non_existing_email@example.com"
				user := &models.User{
					Email: email,
				}

				Expect(user.IsExistingUser()).To(BeFalse())
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users")
	})
})
