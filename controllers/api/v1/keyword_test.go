package apiv1

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/onsi/gomega/gstruct"
	"github.com/tidwall/gjson"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "google-scraper/initializers"
	. "google-scraper/tests"
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
		Context("Given valid credentials", func() {
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
				tokenResponse := KeywordListResponse{
					Id:              gjson.Get(string(responseBody), "data.0.id").Int(),
					Name:            gjson.Get(string(responseBody), "data.0.attributes.name").String(),
					SearchCompleted: gjson.Get(string(responseBody), "data.0.attributes.search_completed").Bool(),
					CreatedAt:       gjson.Get(string(responseBody), "data.0.attributes.created_at").String(),
					Links:           links,
					Meta:            meta,
				}

				Expect(tokenResponse).To(gstruct.MatchAllFields(gstruct.Fields{
					"Id":              BeNumerically("==", keyword.Id),
					"Name":            Equal(keyword.Name),
					"SearchCompleted": BeTrue(),
					"CreatedAt":       Not(BeEmpty()),
					"Links":           Equal(links),
					"Meta":            Equal(meta),
				}))
			})

			Context("given with pagination query string param", func() {
				It("returns only records of the page number 2", func() {
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

			Context("given with keyword query string param", func() {
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

		Context("Given the INVALID credential", func() {
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

	AfterEach(func() {
		TruncateTables("users", "oauth2_clients", "oauth2_tokens")
	})
})
