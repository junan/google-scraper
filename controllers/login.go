package controllers

import (
	"fmt"
	"net/http"

	"google-scraper/forms"

	"github.com/beego/beego/v2/server/web"
)

type Login struct {
	baseController
}

func (c *Login) Get() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Login) Post() {
	loginForm := forms.LoginForm{}
	flash := web.NewFlash()

	err := c.ParseForm(&loginForm)
	if err != nil {
		flash.Error(err.Error())
	}

	_, err = loginForm.Authenticate()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.Data["Form"] = loginForm
		c.setAttributes()
	} else {
		flash.Success("Signed in successfully.")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(http.StatusFound, "/")
	}
}

func (c *Login) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "login/new.html"
	c.Data["Title"] = "Login to your account"
}
