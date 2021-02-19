package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type Dashboard struct {
	baseController
}

func (c *Dashboard) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *Dashboard) Get() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Dashboard) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "dashboard/new.html"
}
