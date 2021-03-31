package apiv1

import (
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

	tokenResponse := c.getTokenResponse(json)

	c.serveJSON(&tokenResponse)
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
