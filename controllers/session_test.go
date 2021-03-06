package controllers_test

import (
	"net/http"
	"net/url"
	"strings"

	_ "google-scraper/initializers"
	. "google-scraper/tests"
	"google-scraper/controllers"

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
				response := MakeAuthenticatedRequest("GET", "/login", nil, nil, &user)
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

					Expect(flash.Data["error"]).To(Equal("Incorrect email or password"))
				})
			})
		})

		Context("given the user is an authenticated user", func() {
			It("does NOT set any flash messages", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				params := url.Values{
					"email":    {email},
					"password": {password},
				}
				body := strings.NewReader(params.Encode())
				response := MakeAuthenticatedRequest("POST", "/login", nil, body, &user)
				flash := GetFlash(response.Cookies())

				Expect(flash).To(BeNil())
			})

			It("redirects to the root path", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				params := url.Values{
					"email":    {email},
					"password": {password},
				}
				body := strings.NewReader(params.Encode())
				response := MakeAuthenticatedRequest("POST", "/login", nil, body, &user)
				path := GetUrlPath(response)

				Expect(path).To(Equal("/"))
			})
		})
	})

	Describe("GET /logout", func() {
		Context("given the user is an authenticated user", func() {
			It("redirects to the login page", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				response := MakeAuthenticatedRequest("GET", "/logout", nil, nil, &user)
				path := GetUrlPath(response)

				Expect(path).To(Equal("/login"))
			})

			It("sets successful logout message", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				response := MakeAuthenticatedRequest("GET", "/logout", nil, nil, &user)
				flash := GetFlash(response.Cookies())

				Expect(flash.Data["success"]).To(Equal("Signed out successfully."))
			})

			It("destroys the user session", func() {
				user := FabricateUser("John", "john@example.com", "secret")
				response := MakeAuthenticatedRequest("GET", "/logout", nil, nil, &user)
				session := GetSession(response.Cookies(), controllers.CurrentUserSession)

				Expect(session).To(BeNil())
			})
		})

		Context("given the user is a guest user", func() {
			It("redirects to the login page", func() {
				response := MakeRequest("GET", "/logout", nil)
				path := GetUrlPath(response)

				Expect(path).To(Equal("/login"))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users")
	})
})
