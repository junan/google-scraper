package apiv1_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	. "google-scraper/helpers"
	. "google-scraper/serializers"
	. "google-scraper/tests"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/tidwall/gjson"
)

var _ = Describe("TokenController", func() {
	Describe("POST /api/v1/token", func() {
		Context("Given the valid credential", func() {
			It("returns 200 status code", func() {
				email := "john@example.com"
				password := "secret"
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				user := FabricateUser("John", email, password)

				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"grant_type":    {"password"},
					"username":      {email},
					"password":      {password},
				}
				body := strings.NewReader(form.Encode())

				response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns correct json response", func() {
				email := "john@example.com"
				password := "secret"
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				user := FabricateUser("John", email, password)
				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"grant_type":    {"password"},
					"username":      {email},
					"password":      {password},
				}
				body := strings.NewReader(form.Encode())
				response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)
				responseBody, err := ioutil.ReadAll(response.Body)
				if err != nil {
					Fail("Reading response body failed: " + err.Error())
				}

				id, err := StringToInt(gjson.Get(string(responseBody), "data.id").String())
				if err != nil {
					Fail("Converting failed: " + err.Error())
				}

				tokenResponse := TokenResponse{
					Id:           id,
					AccessToken:  gjson.Get(string(responseBody), "data.attributes.access_token").String(),
					ExpiresIn:    time.Duration(gjson.Get(string(responseBody), "data.attributes.expires_in").Int()),
					RefreshToken: gjson.Get(string(responseBody), "data.attributes.refresh_token").String(),
					TokenType:    gjson.Get(string(responseBody), "data.attributes.token_type").String(),
				}

				Expect(tokenResponse).To(gstruct.MatchAllFields(gstruct.Fields{
					"Id":           BeNumerically("==", 0),
					"AccessToken":  Not(BeEmpty()),
					"ExpiresIn":    BeNumerically("==", 7200),
					"RefreshToken": Not(BeEmpty()),
					"TokenType":    Equal("Bearer"),
				}))

			})
		})

		Context("Given the INVALID credential", func() {
			Context("Given the user credential is INVALID", func() {
				It("returns 401 unauthorized status code", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {"invalid"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns correct json response", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)
					expectedResponse := `{
						"errors": [
							{
								"detail": "Client authentication failed"
							}
						]
					}`

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {"invalid"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					Expect(string(responseBody)).To(MatchJSON(expectedResponse))
				})
			})

			Context("Given the oauth client credential is INVALID", func() {
				It("returns 401 unauthorized status code", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {"invalid"},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns correct json response", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)
					expectedResponse := `{
						"errors": [
							{
								"detail": "Client authentication failed"
							}
						]
					}`

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {"invalid"},
						"grant_type":    {"password"},
						"username":      {email},
						"password":      {"invalid"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					Expect(string(responseBody)).To(MatchJSON(expectedResponse))
				})
			})

			Context("Given the grant type is INVALID", func() {
				It("returns 401 unauthorized status code", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"grant_type":    {"invalid"},
						"username":      {email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns correct json response", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)
					expectedResponse := `{
						"errors": [
							{
								"detail": "The authorization grant type is not supported by the authorization server"
							}
						]
					}`

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"grant_type":    {"invalid"},
						"username":      {email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					Expect(string(responseBody)).To(MatchJSON(expectedResponse))
				})
			})
		})
	})

	Describe("POST /api/v1/revoke", func() {
		Context("Given the valid credential", func() {
			It("returns 204 no content status code", func() {
				email := "john@example.com"
				password := "secret"
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				token := FabricateOAuthToken(client)
				user := FabricateUser("John", email, password)

				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"token":         {token.GetAccess()},
				}
				body := strings.NewReader(form.Encode())

				response := MakeAuthenticatedRequest("POST", "/api/v1/revoke", body, &user)

				Expect(response.StatusCode).To(Equal(http.StatusNoContent))
			})

			It("returns empty response", func() {
				email := "john@example.com"
				password := "secret"
				client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
				token := FabricateOAuthToken(client)
				user := FabricateUser("John", email, password)

				form := url.Values{
					"client_id":     {client.ID},
					"client_secret": {client.Secret},
					"token":         {token.GetAccess()},
				}
				body := strings.NewReader(form.Encode())

				response := MakeAuthenticatedRequest("POST", "/api/v1/revoke", body, &user)
				responseBody, err := ioutil.ReadAll(response.Body)
				if err != nil {
					Fail("Reading response body failed: " + err.Error())
				}

				Expect(string(responseBody)).To(BeEmpty())
			})
		})

		Context("Given the INVALID credential", func() {
			Context("Given the user token is empty", func() {
				It("returns 401 unauthorized status code", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"token":         {},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/revoke", body, &user)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns correct json response", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					user := FabricateUser("John", email, password)
					expectedResponse := `{
						"errors": [
							{
								"detail": "Token is blank."
							}
						]
					}`

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {client.Secret},
						"token":         {},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/revoke", body, &user)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					Expect(string(responseBody)).To(MatchJSON(expectedResponse))
				})
			})

			Context("Given the oauth client credential is INVALID", func() {
				It("returns 401 unauthorized status code", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					token := FabricateOAuthToken(client)
					user := FabricateUser("John", email, password)

					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {"invalid"},
						"token":         {token.GetAccess()},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/token", body, &user)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns correct json response", func() {
					email := "john@example.com"
					password := "secret"
					client := FabricateOAuthClient(uuid.New().String(), uuid.New().String())
					token := FabricateOAuthToken(client)
					user := FabricateUser("John", email, password)
					expectedResponse := `{
						"errors": [
							{
								"detail": "Client credential invalid."
							}
						]
					}`
					form := url.Values{
						"client_id":     {client.ID},
						"client_secret": {"invalid"},
						"token":         {token.GetAccess()},
					}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/revoke", body, &user)
					responseBody, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Reading response body failed: " + err.Error())
					}

					Expect(string(responseBody)).To(MatchJSON(expectedResponse))
				})
			})
		})
	})
	AfterEach(func() {
		TruncateTables("users", "oauth2_clients", "oauth2_tokens")
	})
})
