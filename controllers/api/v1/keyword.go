package apiv1

import (
	"net/http"

	"github.com/beego/beego/v2/core/logs"

	"google-scraper/forms"
)

type Keyword struct {
	baseAPIController
}

func (c *Keyword) Upload() {
	file, header, err := c.GetFile("file")
	if err != nil {
		logs.Error("Getting file failed: ", err)
	}
	err = forms.PerformSearch(file, header, c.CurrentUser)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusCreated)
}
