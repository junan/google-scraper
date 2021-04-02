package apiv1

import (
	"net/http"

	"google-scraper/forms"
	. "google-scraper/helpers"
	"google-scraper/models"
	"google-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
)

type Keyword struct {
	baseAPIController
}

func (c *Keyword) Upload() {
	tokenInfo, err := oauth.OauthServer.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
		return
	}

	userID, err := StringToInt(tokenInfo.GetUserID())
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
		return
	}

	user, err := models.FindUserById(userID)
	if err != nil {
		c.renderError(err, http.StatusNotFound)
	}

	file, header, err := c.GetFile("file")
	if err != nil {
		logs.Error("Getting file failed: ", err)
	}
	err = forms.PerformSearch(file, header, user)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusCreated)
}
