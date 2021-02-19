package controllers

import (
	"net/http"

	"google-scraper/helpers"
	"google-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const CurrentUserSession = "currentUserSession"

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	web.Controller

	CurrentUser *models.User
	requireAuthenticatedUser bool
	requireGuestUser         bool
}

func (c *baseController) SetCurrentUser(user *models.User) {
	err := c.SetSession(CurrentUserSession, user.Id)
	if err != nil {
		logs.Error("Setting current user failed: ", err)
	}

	c.CurrentUser = user
}

func (c *baseController) GetCurrentUser() (user *models.User) {
	userId := c.GetSession(CurrentUserSession)
	if userId == nil {
		return nil
	}

	user, err := models.FindUserById(userId.(int64))
	if err != nil {
		return nil
	}

	return user
}

func (c *baseController) isGuestUser() bool {
	return c.GetSession(CurrentUserSession) == nil
}

func (c *baseController) isAuthenticatedUser() bool {
	return c.GetCurrentUser() != nil
}

func (c *baseController) Prepare() {
	helpers.SetDataAttributes(&c.Controller)

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}

	c.Data["CurrentUser"] = c.GetCurrentUser()

	if c.requireGuestUser && c.isAuthenticatedUser() {
		c.Ctx.Redirect(http.StatusFound, "/")
	}

	if c.requireAuthenticatedUser && c.isGuestUser() {
		c.Ctx.Redirect(http.StatusFound, "/login")
	}
}
