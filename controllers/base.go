package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"google-scraper/helpers"
)

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	web.Controller
}

func (c *baseController) Prepare() {
	helpers.SetDataAttributes(&c.Controller)

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}
}
