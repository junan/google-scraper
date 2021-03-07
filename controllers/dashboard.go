package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type Dashboard struct {
	baseController
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Dashboard) setAttributes() {
	c.TplName = "dashboard/new.html"
}
