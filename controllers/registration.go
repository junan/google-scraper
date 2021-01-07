package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RegistrationController struct {
	beego.Controller
}

func (c *RegistrationController) Get() {
	c.TplName = "registration.tpl"
}
