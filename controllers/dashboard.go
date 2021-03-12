package controllers

import (
	"google-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type Dashboard struct {
	baseController
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()


	keywords, err := models.FetchKeywords(c.CurrentUser)
	if err != nil {
		logs.Error("Fetching keywords failed: ", err)
	}

	c.Data["Keywords"] = keywords
}

func (c *Dashboard) setAttributes() {
	c.TplName = "dashboard/new.html"
}
