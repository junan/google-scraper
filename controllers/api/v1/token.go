package apiv1

import (
	"net/http"
	"net/http/httptest"

	. "google-scraper/services/oauth"
	"github.com/beego/beego/v2/core/logs"
)

type Token struct {
	baseAPIController
}

func (c *Token) Create() {
	w := httptest.NewRecorder()
	err := OauthServer.HandleTokenRequest(w, c.Ctx.Request)

	re := w.Body.String()

	logs.Info(re)

	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusForbidden)
	}
}

