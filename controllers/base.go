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
	authPolicy  AuthPolicy
	actionName  string
}

type AuthPolicy struct {
	canAccess bool
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
	c.setDefaultAuthPolicy()
	c.setActionName()
	c.setCurrentUser()

	helpers.SetDataAttributes(&c.Controller, c.CurrentUser)

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}

	c.handleRequestAuthorization()
}

func (c *baseController) handleRequestAuthorization() {
	if !c.authPolicy.canAccess {
		c.Ctx.Redirect(http.StatusFound, c.getRedirectPath())
	}
}

func (c *baseController) setDefaultAuthPolicy() {
	// By default user need to log in to access any routes.
	// Override this default policy in the `NestPrepare()` function when necessary(ex: login, registration page)
	c.authPolicy = AuthPolicy{canAccess: c.isAuthenticatedUser()}
}

func (c *baseController) setActionName() {
	_, actionName := c.GetControllerAndAction()
	c.actionName = actionName
}

func (c *baseController) setCurrentUser() {
	c.CurrentUser = c.GetCurrentUser()
}

func (c *baseController) getRedirectPath() string {
	if c.isAuthenticatedUser() {
		return "/"
	} else {
		return "login"
	}
}
