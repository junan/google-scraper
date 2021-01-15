package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RegistrationController struct {
	beego.Controller
}

func (c *RegistrationController) Get() {
	c.TplName = "registration.html"

	//c.Layout = "layouts/authentication.html"
	//c.TplName = "registration/new.html"

	c.Data["Title"] = "Create a new account"
}
