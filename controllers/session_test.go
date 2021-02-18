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

		Context("given the user is an authenticated user", func() {
			It("redirects to the root path", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				response := MakeAuthenticatedRequest("GET", "/login", nil, &user)
				currentPath := GetUrlPath(response)

				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	Describe("POST /login", func() {
		Context("given the user is a guest user", func() {
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

		Context("given the user is an authenticated user", func() {
			It("returns an error", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				params := url.Values{
					"email":    {email},
					"password": {password},
				}
				body := strings.NewReader(params.Encode())
				response := MakeAuthenticatedRequest("POST", "/login", body, &user)
				path := GetUrlPath(response)

				Expect(path).To(Equal("/"))

			})
		})
	})

	AfterEach(func() {
		TruncateTable("users")
	})
})
