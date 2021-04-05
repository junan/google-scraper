package apiv1

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("HealthCheckController", func() {
	Describe("GET /api/v1/health-check", func() {
		It("returns 200 status code", func() {
			//response := MakeRequest("GET", "/api/v1/health-check", nil)
			//
			//Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		//It("returns correct JSON response", func() {
		//	email := "john@example.com"
		//	password := "secret"
		//	client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
		//	FabricateUser("John", email, password)
		//
		//	expectedResponse := `{
		//		"data": {
		//			"type": "health_check",
		//			"id": "0",
		//			"attributes": {
		//				"success": true
		//			}
		//		}
		//	}`
		//
		//	form := url.Values{
		//		"client_id":     {client.ID},
		//		"client_secret": {client.Secret},
		//		"grant_type":    {"password"},
		//		"username":      {email},
		//		"password":      {password},
		//	}
		//	body := strings.NewReader(form.Encode())
		//
		//	response := MakeAuthenticatedRequest("GET", "/api/v1/health-check", nil, body, nil)
		//	responseBody, err := ioutil.ReadAll(response.Body)
		//	if err != nil {
		//		Fail("Reading response body failed: " + err.Error())
		//	}
		//
		//	Expect(string(responseBody)).To(MatchJSON(expectedResponse))
		//})
	})
})
