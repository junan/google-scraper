package apiv1

import (
	"io/ioutil"
	"net/http"

	. "google-scraper/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HealthCheckController", func() {
	Describe("GET /api/v1/health-check", func() {
		It("returns 200 status code", func() {
			response := MakeRequest("GET", "/api/v1/health-check", nil)

			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("returns correct JSON response", func() {
			expectedResponse := `{
				"data": {
					"type": "health_check",
					"id": "0",
					"attributes": {
						"success": true
					}
				}
			}`

			response := MakeRequest("GET", "/api/v1/health-check", nil)
			responseBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				Fail("Reading response body failed: " + err.Error())
			}

			Expect(string(responseBody)).To(MatchJSON(expectedResponse))
		})
	})
})
