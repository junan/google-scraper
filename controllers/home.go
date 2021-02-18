// TODO: remove this temporary controller later

package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type Home struct {
	baseController
}

func (c *Home) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *Home) Get() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Home) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "home/new.html"
}
