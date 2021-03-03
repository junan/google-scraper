package controllers

import (
	. "google-scraper/services/crawler"
	. "google-scraper/constants"
	"github.com/beego/beego/v2/server/web"
)

type Dashboard struct {
	baseController
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()

	Crawl("Buy Domain", GoogleSearchBaseUrl)
}

func (c *Dashboard) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "dashboard/new.html"
}
