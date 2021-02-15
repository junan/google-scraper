package controllers

import (
	"fmt"
	"net/http"

	"google-scraper/forms"

	"github.com/beego/beego/v2/server/web"
)

type Session struct {
	baseController
}

func (c *Session) Get() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Session) Post() {
	sessionForm := forms.SessionForm{}
	flash := web.NewFlash()
	redirectPath := "/"

	err := c.ParseForm(&sessionForm)
	if err != nil {
		flash.Error(err.Error())
	}

	_, err = sessionForm.Authenticate()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		flash.Store(&c.Controller)

		c.Data["Form"] = sessionForm
		redirectPath = "/login"
	} else {
		flash.Success("Signed in successfully.")
		flash.Store(&c.Controller)
	}

	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

func (c *Session) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "session/new.html"
	c.Data["Title"] = "Login to your account"
}
