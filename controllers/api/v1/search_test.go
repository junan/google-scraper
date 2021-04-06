package apiv1

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	_ "google-scraper/initializers"
	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchController", func() {
	Describe("POST /api/v1/search", func() {
		Context("Given valid credentials", func() {
			Context("given the params are valid", func() {
				It("returns 204 status code", func() {
					user := FabricateUser("John", "john@example.com", "secret")
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					filePath := AppRootDir() + "/tests/fixtures/shared/valid_keywords.csv"
					_, body := CreateMultipartFormData(filePath)
					mockResponseFilePath := AppRootDir() + "/tests/fixtures/services/crawler/valid_get_response.html"
					MockCrawling(mockResponseFilePath)

					header := http.Header{
						"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
						"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

					response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)

					Expect(response.StatusCode).To(Equal(http.StatusNoContent))
				})

				It("returns empty response", func() {
					user := FabricateUser("John", "john@example.com", "secret")
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					filePath := AppRootDir() + "/tests/fixtures/shared/valid_keywords.csv"
					_, body := CreateMultipartFormData(filePath)

					header := http.Header{
						"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
						"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

					response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					Expect(string(responseBody)).To(BeEmpty())
				})
			})

			Context("given the params are INVALID", func() {
				Context("given the uploaded file is nil", func() {
					It("returns 422 status code", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						body := CreateEmptyMultipartBody()
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)

						Expect(response.StatusCode).To(Equal(http.StatusUnprocessableEntity))
					})

					It("returns correct json response", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						body := CreateEmptyMultipartBody()
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)
						expectedResponse := `{
							"errors": [
								{
									"detail": "File can't be blank."
								}
							]
						}`

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)
						responseBody, err := ioutil.ReadAll(response.Body)
						if err != nil {
							Fail("Reading response body failed: " + err.Error())
						}

						Expect(string(responseBody)).To(MatchJSON(expectedResponse))
					})
				})

				Context("given the CSV file is wrongly formatted", func() {
					It("returns 422 status code", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/invalid_keyword.csv"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)

						Expect(response.StatusCode).To(Equal(http.StatusUnprocessableEntity))
					})

					It("returns correct json response", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/invalid_keyword.csv"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)
						expectedResponse := `{
							"errors": [
								{
									"detail": "CSV contents are not in correct format."
								}
							]
						}`

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)
						responseBody, err := ioutil.ReadAll(response.Body)
						if err != nil {
							Fail("Reading response body failed: " + err.Error())
						}

						Expect(string(responseBody)).To(MatchJSON(expectedResponse))
					})
				})

				Context("given an invalid file type", func() {
					It("returns 422 status code", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/test.jpeg"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)

						Expect(response.StatusCode).To(Equal(http.StatusUnprocessableEntity))
					})

					It("returns correct json response", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/test.jpeg"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)
						expectedResponse := `{
							"errors": [
								{
									"detail": "Please upload the file in CSV format."
								}
							]
						}`

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)
						responseBody, err := ioutil.ReadAll(response.Body)
						if err != nil {
							Fail("Reading response body failed: " + err.Error())
						}

						Expect(string(responseBody)).To(MatchJSON(expectedResponse))
					})
				})

				Context("given the file size is more than 5 megabytes", func() {
					It("returns 422 status code", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/big_file.pdf"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)

						Expect(response.StatusCode).To(Equal(http.StatusUnprocessableEntity))
					})

					It("returns correct json response", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/big_file.pdf"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)
						expectedResponse := `{
							"errors": [
								{
									"detail": "File size can't be more than 5 MB."
								}
							]
						}`

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)
						responseBody, err := ioutil.ReadAll(response.Body)
						if err != nil {
							Fail("Reading response body failed: " + err.Error())
						}

						Expect(string(responseBody)).To(MatchJSON(expectedResponse))
					})
				})
				Context("given the upload file has NO keywords", func() {
					It("returns 422 status code", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/empty_keyword.csv"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)

						Expect(response.StatusCode).To(Equal(http.StatusUnprocessableEntity))
					})

					It("returns correct json response", func() {
						user := FabricateUser("John", "john@example.com", "secret")
						filePath := AppRootDir() + "/tests/fixtures/shared/empty_keyword.csv"
						_, body := CreateMultipartFormData(filePath)
						client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
						oauthToken := FabricateOAuthToken(client, user.Id)
						expectedResponse := `{
							"errors": [
								{
									"detail": "Keywords count can't be more than 1000 or less than 1."
								}
							]
						}`

						header := http.Header{
							"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())},
							"Content-Type":  {"multipart/form-data; boundary=multipart-boundary"}}

						response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, body, nil)
						responseBody, err := ioutil.ReadAll(response.Body)
						if err != nil {
							Fail("Reading response body failed: " + err.Error())
						}

						Expect(string(responseBody)).To(MatchJSON(expectedResponse))
					})
				})
			})
		})
	})

	Context("Given the INVALID credential", func() {
		It("returns 401 status code", func() {
			header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", "invalid")}}

			response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, nil, nil)

			Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
		})

		It("returns correct JSON response", func() {
			header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", "invalid")}}
			expectedResponse := `{
				"errors": [
					{
						"detail": "Client authentication failed"
					}
				]
			}`

			response := MakeAuthenticatedRequest("POST", "/api/v1/search", header, nil, nil)

			responseBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				Fail("Reading response body failed: " + err.Error())
			}

			Expect(string(responseBody)).To(MatchJSON(expectedResponse))
		})
	})
	AfterEach(func() {
		TruncateTables("users", "oauth2_clients", "oauth2_tokens")
	})
})
