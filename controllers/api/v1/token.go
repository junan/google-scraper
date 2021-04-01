package apiv1

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	"google-scraper/serializers"
	. "google-scraper/services/oauth"

	"github.com/tidwall/gjson"
)

type Token struct {
	baseAPIController
}

func (c *Token) Create() {
	writer := httptest.NewRecorder()
	err := OauthServer.HandleTokenRequest(writer, c.Ctx.Request)
	if err != nil {
		c.renderError(err, http.StatusUnauthorized)
		return
	}

	json := writer.Body.String()
	if writer.Code != 200 {
		errorMessage := gjson.Get(json, "error_description").String()
		c.renderError(errors.New(errorMessage), writer.Code)
		return
	}

	tokenResponse := serializers.GetTokenResponse(json)

	c.serveJSON(&tokenResponse)
}

func (c *Token) Revoke() {
	c.authenticateClient()
	token := c.GetString("token")
	if token == "" {
		err := errors.New("Token is blank.")
		c.renderError(err, http.StatusUnauthorized)
		return
	}

	// Remove the token from database
	err := TokenStore.RemoveByAccess(context.TODO(), token)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
		return
	}

	// Response with 204 status code and no body
	c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
}

func (c *Token) getTokenResponse(json string) serializers.TokenResponse {
	response := serializers.TokenResponse{
		AccessToken:  gjson.Get(json, "access_token").String(),
		ExpiresIn:    gjson.Get(json, "expires_in").Int(),
		RefreshToken: gjson.Get(json, "refresh_token").String(),
		TokenType:    gjson.Get(json, "token_type").String(),
	}

	return response
}

func (c *Token) authenticateClient()  {
	clientID := c.GetString("client_id")
	clientSecret := c.GetString("client_secret")
	client, err := ClientStore.GetByID(context.TODO(), clientID)
	if err != nil {
		err = errors.New("Client credential invalid.")
		c.renderError(err, http.StatusUnauthorized)
		return
	}

	if client.GetSecret() != clientSecret {
		err = errors.New("Client credential invalid.")
		c.renderError(err, http.StatusUnauthorized)
		return
	}
}
