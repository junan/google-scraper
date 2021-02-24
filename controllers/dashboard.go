package controllers

import (
	"github.com/beego/beego/v2/server/web"
	. "google-scraper/scraper"
)

type Dashboard struct {
	baseController
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()

	Crawl("Ruby on Rails")
}

func (c *Dashboard) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "dashboard/new.html"
}
