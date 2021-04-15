package apiv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "google-scraper/initializers"
	"google-scraper/serializers"
	. "google-scraper/tests"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/tidwall/gjson"
)

type KeywordListResponse struct {
	Id              int64
	Name            string
	SearchCompleted bool
	CreatedAt       string
	Meta            gjson.Result
	Links           gjson.Result
}

var _ = Describe("KeywordController", func() {
	Describe("GET /api/v1/keywords", func() {
		Context("given valid credentials", func() {
			It("returns 200 status code", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				FabricateKeyword("Buy domain", false, &user)
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				oauthToken := FabricateOAuthToken(client, user.Id)
				header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", header, nil, nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns correct JSON response", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				keyword := FabricateKeyword("Buy domain", true, &user)
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				oauthToken := FabricateOAuthToken(client, user.Id)
				header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", header, nil, nil)
				responseBody, err := ioutil.ReadAll(response.Body)
				if err != nil {
					Fail("Reading response body failed: " + err.Error())
				}

				links := gjson.Get(string(responseBody), "links")
				meta := gjson.Get(string(responseBody), "meta")
				keywordListResponse := KeywordListResponse{
					Id:              gjson.Get(string(responseBody), "data.0.id").Int(),
					Name:            gjson.Get(string(responseBody), "data.0.attributes.name").String(),
					SearchCompleted: gjson.Get(string(responseBody), "data.0.attributes.search_completed").Bool(),
					CreatedAt:       gjson.Get(string(responseBody), "data.0.attributes.created_at").String(),
					Links:           links,
					Meta:            meta,
				}

				Expect(keywordListResponse).To(gstruct.MatchAllFields(gstruct.Fields{
					"Id":              BeNumerically("==", keyword.Id),
					"Name":            Equal(keyword.Name),
					"SearchCompleted": BeTrue(),
					"CreatedAt":       Not(BeEmpty()),
					"Links":           Equal(links),
					"Meta":            Equal(meta),
				}))
			})

			Context("given pagination query string param", func() {
				It("returns only records of the requested page", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					for i := 1; i <= 11; i++ {
						FabricateKeyword(fmt.Sprintf("Buy domain %d", i), false, &user)
					}

					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

					response := MakeAuthenticatedRequest("GET", "/api/v1/keywords?p=2", header, nil, nil)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					keywords := gjson.Get(string(responseBody), "data").Get("#")

					Expect(int(keywords.Num)).To(Equal(1))
				})
			})

			Context("given keyword query string param", func() {
				It("returns only matched keywords", func() {
					keyword := "Buy_domain"
					user := FabricateUser("John", "john@example.com", "secret")

					// Created three keyword with different name
					FabricateKeyword(keyword, false, &user)
					FabricateKeyword("Purchase domain", false, &user)
					FabricateKeyword("Buy bike", false, &user)

					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/keywords?keyword=%s", keyword), header, nil, nil)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					keywords := gjson.Get(string(responseBody), "data").Get("#")
					returnedKeyword := gjson.Get(string(responseBody), "data.0.attributes.name").String()

					Expect(int(keywords.Num)).To(Equal(1))
					Expect(returnedKeyword).To(Equal(keyword))
				})
			})
		})

		Context("given the INVALID credential", func() {
			It("returns 401 status code", func() {
				header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", "invalid")}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", header, nil, nil)

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

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", header, nil, nil)

				responseBody, err := ioutil.ReadAll(response.Body)
				if err != nil {
					Fail("Reading response body failed: " + err.Error())
				}

				Expect(string(responseBody)).To(MatchJSON(expectedResponse))
			})
		})
	})

	Describe("GET /api/v1/keywords/:id", func() {
		Context("Given valid credentials", func() {
			Context("Given existing keyword ID", func() {
				It("returns 200 status code", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					keyword := FabricateKeyword("Buy domain", false, &user)
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/keywords?keyword=%d", keyword.Id), header, nil, nil)

					Expect(response.StatusCode).To(Equal(http.StatusOK))
				})

				It("returns correct JSON response", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					keyword := FabricateKeyword("Buy domain", true, &user)
					FabricateSearchResult(&keyword)
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/keywords/%d", keyword.Id), header, nil, nil)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					resultUrlsString := gjson.Get(string(responseBody), "data.attributes.results_url").String()
					adwordUrlsString := gjson.Get(string(responseBody), "data.attributes.top_ad_word_advertisers_urls").String()
					var resultURLs []string
					var adwordURLs []string
					err = json.Unmarshal([]byte(resultUrlsString), &resultURLs)
					if err != nil {
						Fail("Unmarshal failed: " + err.Error())
					}

					err = json.Unmarshal([]byte(adwordUrlsString), &adwordURLs)
					if err != nil {
						Fail("Unmarshal failed: " + err.Error())
					}

					keywordResponse := serializers.KeywordResponse{
						Id:                          gjson.Get(string(responseBody), "data.id").Int(),
						Name:                        gjson.Get(string(responseBody), "data.attributes.name").String(),
						CreatedAt:                   gjson.Get(string(responseBody), "data.attributes.created_at").String(),
						TopAdWordAdvertisersCount:   int(gjson.Get(string(responseBody), "data.attributes.top_ad_word_advertisers_count").Int()),
						TotalAdWordAdvertisersCount: int(gjson.Get(string(responseBody), "data.attributes.total_ad_word_advertisers_count").Int()),
						TotalLinksCount:             int(gjson.Get(string(responseBody), "data.attributes.total_links_count").Int()),
						ResultsCount:                int(gjson.Get(string(responseBody), "data.attributes.results_count").Int()),
						Html:                        gjson.Get(string(responseBody), "data.attributes.html").String(),
						TopAdWordAdvertisersUrls:    adwordURLs,
						ResultsUrls:                 resultURLs,
					}

					Expect(keywordResponse).To(gstruct.MatchAllFields(gstruct.Fields{
						"Id":                          BeNumerically("==", keyword.Id),
						"Name":                        Equal(keyword.Name),
						"CreatedAt":                   Not(BeEmpty()),
						"TopAdWordAdvertisersCount":   BeNumerically("==", 2),
						"TotalAdWordAdvertisersCount": BeNumerically("==", 3),
						"ResultsCount":                BeNumerically("==", 2),
						"TotalLinksCount":             BeNumerically("==", 20),
						"Html":                        Not(BeEmpty()),
						"TopAdWordAdvertisersUrls":    Not(BeEmpty()),
						"ResultsUrls":                 Not(BeEmpty()),
					}))
				})
			})

			Context("Given NON existing keyword ID", func() {
				It("returns 404 status code", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

					response := MakeAuthenticatedRequest("GET", "/api/v1/keywords/2000", header, nil, nil)

					Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				})

				It("returns correct JSON response", func() {
					email := "john@example.com"
					password := "secret"
					user := FabricateUser("John", email, password)
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					oauthToken := FabricateOAuthToken(client, user.Id)
					header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}
					expectedResponse := `{
						"errors": [
							{
								"detail": "Keyword not found."
							}
						]
					}`

					response := MakeAuthenticatedRequest("GET", "/api/v1/keywords/2000", header, nil, nil)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					Expect(string(responseBody)).To(MatchJSON(expectedResponse))

				})
			})
		})

		Context("Given the INVALID credential", func() {
			It("returns 401 status code", func() {
				header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", "invalid")}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords/1", header, nil, nil)

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

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords/1", header, nil, nil)

				responseBody, err := ioutil.ReadAll(response.Body)
				if err != nil {
					Fail("Reading response body failed: " + err.Error())
				}

				Expect(string(responseBody)).To(MatchJSON(expectedResponse))
			})
		})
	})

	AfterEach(func() {
		TruncateTables("users", "oauth2_clients", "oauth2_tokens")
	})
})
