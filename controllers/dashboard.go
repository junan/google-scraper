package controllers

import (
	"google-scraper/models"
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

	searchString := "Buy Domain"

	search := &models.Search{
		Keyword: searchString,
		User: c.CurrentUser,
	}

	_, _ = models.CreateResult(search)

	Crawl(searchString, GoogleSearchBaseUrl, search)
}

func (c *Dashboard) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "dashboard/new.html"
}
