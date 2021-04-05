package controllers_test

import (
	"net/http"

	_ "google-scraper/initializers"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DashboardController", func() {
	Describe("GET /", func() {
		Context("given the user is a guest user", func() {
			It("redirects to the login page", func() {
				response := MakeRequest("GET", "/", nil)
				currentPath := GetUrlPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/login"))
			})
		})

		Context("given the user is an authenticated user", func() {
			It("returns 200 status code", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				response := MakeAuthenticatedRequest("GET", "/", nil, nil, &user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users")
	})
})
