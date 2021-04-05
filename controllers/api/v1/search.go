package apiv1

import (
	"net/http"

	"google-scraper/forms"

	"github.com/beego/beego/v2/core/logs"
)

type Search struct {
	baseAPIController
}

func (c *Search) Create() {
	file, header, err := c.GetFile("file")
	if err != nil {
		logs.Error("Getting file failed: ", err)
	}
	err = forms.PerformSearch(file, header, c.CurrentUser)
	if err != nil {
		c.renderError(err, http.StatusUnprocessableEntity)
	}

	c.Ctx.ResponseWriter.WriteHeader(http.StatusCreated)
}
