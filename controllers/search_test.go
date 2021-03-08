package controllers_test

import (
	. "google-scraper/tests/fabricators"
	. "google-scraper/tests/testing_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchController", func() {
	Describe("POST", func() {
		Context("given the user is an authenticated user", func() {
			Context("given the user is uploaded a valid CSV file", func() {
				It("sets flash success message", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					body := CreateMultipartFormData("valid_keywords.csv")

					response := MakeAuthenticatedRequest("POST", "/search", body, &user)

					flash := GetFlash(response.Cookies())

					Expect(flash.Data["success"]).To(Equal("Your csv file has been uploaded successfully"))
				})

				It("redirects to the root path", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					body := CreateMultipartFormData("valid_keywords.csv")

					response := MakeAuthenticatedRequest("POST", "/search", body, &user)

					path := GetUrlPath(response)

					Expect(path).To(Equal("/"))
				})
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords", "search_results")
	})
})
