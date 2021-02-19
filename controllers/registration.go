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

func (c *Registration) NestPrepare() {
	c.setRegistrationPolicy()
}

func (c *Registration) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()
}

func (c *Registration) Create() {
	registrationForm := forms.RegistrationForm{}
	flash := web.NewFlash()
	redirectPath := "/"

	err := c.ParseForm(&registrationForm)
	if err != nil {
		flash.Error(err.Error())
	}

	user, err := registrationForm.Save()
	if err != nil {
		flash.Error(fmt.Sprint(err))
		redirectPath = "/register"
	} else {
		c.SetCurrentUser(user)
		flash.Success("Account has been created successfully")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

func (c *Registration) setRegistrationPolicy() {
	_, actionName := c.GetControllerAndAction()
	p := Policy{redirectPath: "/login"}

	if actionName == "New" || actionName == "Create" {
		p.requireAuthorization = c.isAuthenticatedUser()
		p.redirectPath = "/"
	} else {
		p.requireAuthorization = true
	}

	c.requestPolicy[actionName] = p
}

func (c *Registration) setAttributes() {
	c.Layout = "layouts/authentication.html"
	c.TplName = "registration/new.html"
	c.Data["Title"] = "Create your account"
}
