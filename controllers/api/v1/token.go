package apiv1

import (
	"net/http"

	. "google-scraper/services/oauth"
)

type Token struct {
	baseAPIController
}

func (c *Token) Create() {
	err := OauthServer.HandleTokenRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusForbidden)
	}
}

