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

func (c *Token) NestPrepare() {
	c.setTokenPolicy()
}

func (c *Token) Create() {
	writer := httptest.NewRecorder()
	err := OauthServer.HandleTokenRequest(writer, c.Ctx.Request)
	if err != nil {
		c.renderError(err, http.StatusUnauthorized)
	}

	json := writer.Body.String()
	if writer.Code != 200 {
		errorMessage := gjson.Get(json, "error_description").String()
		c.renderError(errors.New(errorMessage), writer.Code)
	}

	tokenResponse := serializers.GetTokenResponse(json)

	c.serveJSON(&tokenResponse)
}

func (c *Token) Revoke() {
	token := c.GetString("access_token")
	// Remove the token from database
	err := TokenStore.RemoveByAccess(context.TODO(), token)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	// Response with 204 status code and no body
	c.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
}

func (c *Token) setTokenPolicy() {
	if c.actionName == "Create" {
		c.authPolicy.requireTokenValidation = false
	} else if c.actionName == "Revoke" {
		c.authPolicy.requireClientCredentialValidation = true
	}
}
