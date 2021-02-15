package controllers

import (
	"fmt"
	"net/http"

	"google-scraper/forms"

	"github.com/beego/beego/v2/server/web"
)

type Registration struct {
	baseController
}

func (c *Registration) Get() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Registration) Post() {
	registrationForm := forms.RegistrationForm{}
	flash := web.NewFlash()

	err := c.ParseForm(&registrationForm)
	if err != nil {
		flash.Error(err.Error())
	}

	_, err = registrationForm.Save()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.Data["Form"] = registrationForm
		c.setAttributes()
	} else {
		flash.Success("Account has been created successfully")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(http.StatusFound, "/register")
	}
}

func (c *Registration) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "registration/new.html"
	c.Data["Title"] = "Create your account"
}
