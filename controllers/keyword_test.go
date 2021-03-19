package controllers_test

import (
	"fmt"
	"net/http"

	_ "google-scraper/initializers"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordController", func() {
	Describe("GET /keyword/:id", func() {
		Context("given the user is an authenticated user", func() {
			Context("given the keyword is belong to the user", func() {
				It("returns 200 status code", func() {
					user := FabricateUser("John", "john@example.com", "secret")
					keyword := FabricateKeyword("Buy domain", false, &user)
					url := fmt.Sprintf("/keyword/%d", keyword.Id)

					response := MakeAuthenticatedRequest("GET", url, nil, &user)

					Expect(response.StatusCode).To(Equal(http.StatusOK))
				})
			})

			Context("given the keyword is NOT belong to the user", func() {
				It("redirects to the root path", func() {
					user1 := FabricateUser("John", "john@example.com", "secret")
					FabricateKeyword("Buy domain", false, &user1)
					user2 := FabricateUser("Mike", "mike@example.com", "secret")
					keyword2 := FabricateKeyword("Buy bike", false, &user2)
					url := fmt.Sprintf("/keyword/%d", keyword2.Id)

					response := MakeAuthenticatedRequest("GET", url, nil, &user1)
					currentPath := GetUrlPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/"))
				})
			})

			Context("given the keyword does NOT exist in the database", func() {
				It("redirects to the root path", func() {
					user := FabricateUser("John", "john@example.com", "secret")
					FabricateKeyword("Buy domain", false, &user)
					url := fmt.Sprintf("/keyword/%d", 1000)

					response := MakeAuthenticatedRequest("GET", url, nil, &user)
					currentPath := GetUrlPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/"))
				})
			})
		})

		Context("given the user is a guest user", func() {
			It("redirects to the login page", func() {
				url := fmt.Sprintf("/keyword/%d", 1)

				response := MakeRequest("GET", url, nil)
				currentPath := GetUrlPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/login"))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords")
	})
})
