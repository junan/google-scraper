package apiv1

import (
	"errors"
	"net/http/httptest"

	"github.com/tidwall/gjson"

	"google-scraper/serializers"
	. "google-scraper/services/oauth"
)

type Token struct {
	baseAPIController
}

func (c *Token) Create() {
	writer := httptest.NewRecorder()
	err := OauthServer.HandleTokenRequest(writer, c.Ctx.Request)
	if err != nil {
		c.renderError(err)
		return
	}

	json := writer.Body.String()

	if writer.Code != 200 {
		errorMessage := gjson.Get(json, "error_description").String()
		c.renderError(errors.New(errorMessage))
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
