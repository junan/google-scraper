package controllers_test

import (
	"net/http"
	"net/url"
	"strings"

	. "google-scraper/tests/testing_helpers"
	. "google-scraper/tests/fabricators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegistrationController", func() {
	AfterEach(func() {
		TruncateTable("users")
	})

	Describe("GET", func() {
		Context("given the user is a guest user", func() {
			It("returns 200 status code", func() {
				response := MakeRequest("GET", "/register", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("given the user is an authenticated user", func() {
			It("redirects to the root path", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				response := MakeAuthenticatedRequest("GET", "/register", nil, &user)
				currentPath := GetUrlPath(response)

				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	Describe("POST", func() {
		// TODO: Will add a test case for redirecting to a success page once I add it
		Context("given valid params", func() {
			It("sets flash success message", func() {
				registrationForm := url.Values{
					"name":     {"John"},
					"email":    {"john@example.com"},
					"password": {"secret"},
				}
				body := strings.NewReader(registrationForm.Encode())
				response := MakeRequest("POST", "/register", body)
				flash := GetFlash(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("Account has been created successfully"))
			})
		})

		Context("given invalid params", func() {
			It("sets flash error message", func() {
				form := url.Values{
					"name":     {""},
					"email":    {"john@example.com"},
					"password": {"secret"},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/register", body)
				flash := GetFlash(response.Cookies())

				Expect(flash.Data["error"]).To(Equal("Name can not be empty"))
			})

			It("re-renders the registration page", func() {
				form := url.Values{
					"name":     {""},
					"email":    {"john@example.com"},
					"password": {"secret"},
				}
				body := strings.NewReader(form.Encode())
				response := MakeRequest("POST", "/register", body)
				path := GetUrlPath(response)

				Expect(path).To(Equal("/register"))
			})
		})
	})
})
