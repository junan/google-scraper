package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type RegistrationController struct {
	beego.Controller
}

func (c *RegistrationController) Get() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "registrations/new.html"

	c.Data["Title"] = "Create your account"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["FlashMessage"] = "shared/flash.html"
}
