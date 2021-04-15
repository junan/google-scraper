package serializers

import (
	"time"
	"github.com/tidwall/gjson"
)

type TokenResponse struct {
	Id           int64  `jsonapi:"primary,token"`
	AccessToken  string `jsonapi:"attr,access_token"`
	ExpiresIn    time.Duration  `jsonapi:"attr,expires_in"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
	TokenType    string `jsonapi:"attr,token_type"`
}

func GetTokenResponse(json string) TokenResponse {
	response := TokenResponse{
		AccessToken:  gjson.Get(json, "access_token").String(),
		ExpiresIn:    time.Duration(gjson.Get(json, "expires_in").Int()),
		RefreshToken: gjson.Get(json, "refresh_token").String(),
		TokenType:    gjson.Get(json, "token_type").String(),
	}

	return response
}
