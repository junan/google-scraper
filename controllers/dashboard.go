package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type Dashboard struct {
	baseController
}

func (c *Dashboard) NestPrepare() {
	c.setDashboardPolicy()
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Dashboard) setDashboardPolicy() {
	_, actionName := c.GetControllerAndAction()
	p := Policy{redirectPath: "/login"}

	if actionName == "New"  {
		p.requireAuthorization = c.isGuestUser()
	} else {
		p.requireAuthorization = true
	}

	c.requestPolicy[actionName] = p
}

func (c *Dashboard) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "dashboard/new.html"
}
