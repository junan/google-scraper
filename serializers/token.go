package serializers

type TokenResponse struct {
	Id           int64    `jsonapi:"primary,token"`
	AccessToken  string `jsonapi:"attr,access_token"`
	ExpiresIn    int64  `jsonapi:"attr,expires_in"`
	RefreshToken string `jsonapi:"attr,refresh_token"`
	TokenType    string `jsonapi:"attr,token_type"`
}
