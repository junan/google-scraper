package controllers_test

import (
	"net/http"
	"net/url"
	"strings"

	_ "google-scraper/initializers"
	. "google-scraper/tests/fabricators"
	. "google-scraper/tests/testing_helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionController", func() {
	Describe("GET /login", func() {
		Context("given the user is a guest user", func() {
			It("returns 200 status code", func() {
				response := MakeRequest("GET", "/login", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})

	Describe("POST /login", func() {
		Context("given the session params are valid", func() {
			It("redirects to the root path", func() {
				email := "john@example.com"
				password := "secret"
				FabricateUser("John", email, password)
				params := url.Values{
					"email":    {email},
					"password": {password},
				}
				body := strings.NewReader(params.Encode())

				response := MakeRequest("POST", "/login", body)
				path := GetUrlPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(path).To(Equal("/"))
			})

			It("sets the flash success message", func() {
				email := "john@example.com"
				password := "secret"
				FabricateUser("John", email, password)
				params := url.Values{
					"email":    {email},
					"password": {password},
				}

				body := strings.NewReader(params.Encode())
				response := MakeRequest("POST", "/login", body)
				flash := GetFlash(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("Signed in successfully."))
			})
		})

		Context("given the session params are INVALID", func() {
			It("redirects to the login path", func() {
				email := "invalid-email"
				password := "secret"
				FabricateUser("John", email, password)
				params := url.Values{
					"email":    {email},
					"password": {password},
				}
				body := strings.NewReader(params.Encode())

				response := MakeRequest("POST", "/login", body)
				path := GetUrlPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(path).To(Equal("/login"))
			})

			It("sets the flash error message", func() {
				email := "invalid-email"
				password := "secret"
				FabricateUser("John", email, password)
				params := url.Values{
					"email":    {email},
					"password": {password},
				}
				body := strings.NewReader(params.Encode())

				response := MakeRequest("POST", "/login", body)
				flash := GetFlash(response.Cookies())

				Expect(flash.Data["error"]).To(Equal("Email must be a valid email address"))
			})
		})
	})

	AfterEach(func() {
		TruncateTable("users")
	})
})
