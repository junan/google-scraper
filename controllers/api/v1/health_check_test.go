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

var _ = Describe("HealthCheckController", func() {
	Describe("GET /api/v1/health-check", func() {
		Context("Given the valid credential", func() {
			It("returns 200 status code", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				oauthToken := FabricateOAuthToken(client, user.Id)
				header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/health-check", header, nil, nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns correct JSON response", func() {
				email := "john@example.com"
				password := "secret"
				user := FabricateUser("John", email, password)
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				oauthToken := FabricateOAuthToken(client, user.Id)
				header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", oauthToken.GetAccess())}}
				expectedResponse := `{
				"data": {
					"type": "health_check",
					"id": "0",
					"attributes": {
						"success": true
					}
				}
			}`

				response := MakeAuthenticatedRequest("GET", "/api/v1/health-check", header, nil, nil)

				responseBody, err := ioutil.ReadAll(response.Body)
				if err != nil {
					Fail("Reading response body failed: " + err.Error())
				}

				Expect(string(responseBody)).To(MatchJSON(expectedResponse))
			})
		})

		Context("Given the INVALID credential", func() {
			It("returns 401 status code", func() {
				header := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", "invalid")}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/health-check", header, nil, nil)

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

				response := MakeAuthenticatedRequest("GET", "/api/v1/health-check", header, nil, nil)

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
