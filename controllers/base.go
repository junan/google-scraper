package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	web.Controller
}

func (c *baseController) Prepare() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["FlashMessage"] = "shared/flash.html"

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}
}
