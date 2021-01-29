package controllers

import (
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"google-scraper/forms"
)

type CommonTemplate struct {
	web.Controller
}

func (c *CommonTemplate) Prepare() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "registrations/new.html"

	c.Data["Title"] = "Create your account"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["FlashMessage"] = "shared/flash.html"
}

type RegistrationController struct {
	CommonTemplate
}

func (c *RegistrationController) Get() {
	web.ReadFromRequest(&c.Controller)
}

func (c *RegistrationController) Post() {
	form := forms.RegistrationForm{}
	flash := web.NewFlash()

	err := c.ParseForm(&form)
	if err != nil {
		flash.Error(err.Error())
	}

	err = form.Save()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.Data["Form"] = form
	} else {
		flash.Success("Account created successfully")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(http.StatusFound, "/register")
	}
}
