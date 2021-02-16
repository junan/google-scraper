package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type Test struct {
	baseController
}

func (c *Test) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *Test) Get() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Test) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "test/new.html"
}
