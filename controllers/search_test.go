package controllers_test

import (
	"net/http"

	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchController", func() {
	Describe("POST", func() {
		Context("given the user is an authenticated user", func() {
			Context("given the form params are valid", func() {
				Context("given the uploaded file is a valid CSV file", func() {
					It("sets flash success message", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/valid_keywords.csv"
						header := http.Header{"Content-Type": {"multipart/form-data; boundary=multipart-boundary"}}
						_, body := CreateMultipartFormData(filePath)
						mockResponseFilePath := AppRootDir() + "/tests/fixtures/services/crawler/valid_get_response.html"
						MockCrawling(mockResponseFilePath)

						response := MakeAuthenticatedRequest("POST", "/search", header, body, &user)

						flash := GetFlash(response.Cookies())

						Expect(flash.Data["success"]).To(Equal("Your csv file has been uploaded successfully"))
					})

					It("redirects to the root path", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/valid_keywords.csv"
						header := http.Header{"Content-Type": {"multipart/form-data; boundary=multipart-boundary"}}
						_, body := CreateMultipartFormData(filePath)
						mockResponseFilePath := AppRootDir() + "/tests/fixtures/services/crawler/valid_get_response.html"
						MockCrawling(mockResponseFilePath)

						response := MakeAuthenticatedRequest("POST", "/search", header, body, &user)

						path := GetUrlPath(response)

						Expect(path).To(Equal("/"))
					})
				})
			})

			Context("given the form params are INVALID", func() {
				Context("given the uploaded file is nil", func() {
					It("sets flash error message", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						header := http.Header{"Content-Type": {"multipart/form-data; boundary=multipart-boundary"}}
						body := CreateEmptyMultipartBody()

						response := MakeAuthenticatedRequest("POST", "/search", header, body, &user)

						flash := GetFlash(response.Cookies())

						Expect(flash.Data["error"]).To(Equal("File can't be blank."))
					})
				})

				Context("given the file keywords are empty", func() {
					It("sets flash error message", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/empty_keyword.csv"
						header := http.Header{"Content-Type": {"multipart/form-data; boundary=multipart-boundary"}}
						_, body := CreateMultipartFormData(filePath)

						response := MakeAuthenticatedRequest("POST", "/search", header, body, &user)

						flash := GetFlash(response.Cookies())

						Expect(flash.Data["error"]).To(Equal("Keywords count can't be more than 1000 or less than 1."))
					})
				})

				Context("given the CSV file is wrong formatted", func() {
					It("sets flash error message", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/invalid_keyword.csv"
						header := http.Header{"Content-Type": {"multipart/form-data; boundary=multipart-boundary"}}
						_, body := CreateMultipartFormData(filePath)

						response := MakeAuthenticatedRequest("POST", "/search", header, body, &user)

						flash := GetFlash(response.Cookies())

						Expect(flash.Data["error"]).To(Equal("CSV contents are not in correct format."))
					})
				})

				Context("given the file is an image", func() {
					It("sets flash error message", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/test.jpeg"
						header := http.Header{"Content-Type": {"multipart/form-data; boundary=multipart-boundary"}}
						_, body := CreateMultipartFormData(filePath)

						response := MakeAuthenticatedRequest("POST", "/search", header, body, &user)

						flash := GetFlash(response.Cookies())

						Expect(flash.Data["error"]).To(Equal("Please upload the file in CSV format."))
					})
				})

				Context("given the file size is more than 5 megabytes", func() {
					It("sets flash error message", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/big_file.pdf"
						header := http.Header{"Content-Type": {"multipart/form-data; boundary=multipart-boundary"}}
						_, body := CreateMultipartFormData(filePath)

						response := MakeAuthenticatedRequest("POST", "/search", header, body, &user)

						flash := GetFlash(response.Cookies())

						Expect(flash.Data["error"]).To(Equal("File size can't be more than 5 MB."))
					})
				})
			})
		})

		Context("given the user is a guest user", func() {
			It("redirects to the login path", func() {
				body := CreateEmptyMultipartBody()

				response := MakeRequest("POST", "/search", body)

				path := GetUrlPath(response)

				Expect(path).To(Equal("/login"))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "keywords", "search_results")
	})
})
