package serializers

import (
	"time"
	"github.com/tidwall/gjson"
)

type KeywordListResponse struct {
	Id           int64  `jsonapi:"primary,keyword"`
	Name  string `jsonapi:"attr,name"`
	SearchCompleted    bool  `jsonapi:"attr,search_completed"`
	CreatedAt string `jsonapi:"attr,created_at"`
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
