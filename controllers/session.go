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

func (c *Session) NestPrepare() {
	c.setSessionPolicy()
}

func (c *Session) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Session) Create() {
	sessionForm := forms.SessionForm{}
	flash := web.NewFlash()
	redirectPath := "/"

	err := c.ParseForm(&sessionForm)
	if err != nil {
		flash.Error(err.Error())
	}

	user, err := sessionForm.Authenticate()
	if err != nil {
		flash.Error(fmt.Sprint(err))

		redirectPath = "/login"
	} else {
		c.Controller.Data["CurrentUser"] = user
		c.SetCurrentUser(user)
		flash.Success("Signed in successfully.")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

func (c *Session) Delete() {
	flash := web.NewFlash()
	redirectPath := "/"

	err := c.logout()
	if err != nil {
		flash.Error("Failed to sign out")
	} else {
		flash.Success("Signed out successfully.")
		redirectPath = "/login"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}

func (c *Session) logout() error {
	err := c.DelSession(CurrentUserSession)
	if err == nil {
		c.CurrentUser = nil
	}

	return err
}

func (c *Session) setSessionPolicy() {
	if c.actionName == "New" || c.actionName == "Create" {
		c.authPolicy.canAccess = c.isGuestUser()
	}
}

func (c *Session) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "session/new.html"
	c.Data["Title"] = "Login to your account"
}
