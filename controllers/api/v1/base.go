package apiv1

import (
	"context"
	"errors"
	"net/http"

	. "google-scraper/helpers"
	"google-scraper/models"
	"google-scraper/services/oauth"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/jsonapi"
)

type baseAPIController struct {
	web.Controller

	CurrentUser *models.User
	authPolicy  AuthPolicy
	UserID      string
}

type AuthPolicy struct {
	requireTokenValidation            bool
	requireClientCredentialValidation bool
}

type NestPreparer interface {
	NestPrepare()
}

func (c *baseAPIController) Prepare() {
	c.setDefaultAuthPolicy()
	c.validateAuthorization()
	c.setCurrentUser()
}

func (c *baseAPIController) setDefaultAuthPolicy() {
	// By default it will validate API token  and will not validate client credential for all routes.
	// Override this default policy in the `NestPrepare()` function when necessary(ex: generating api token)
	c.authPolicy = AuthPolicy{requireTokenValidation: true, requireClientCredentialValidation: false}
}

func (c *baseAPIController) validateAuthorization() {
	if c.authPolicy.requireClientCredentialValidation {
		c.validateClientCredential()
	}

	if c.authPolicy.requireTokenValidation {
		c.validateTokenCredential()
	}
}

func (c *baseAPIController) setCurrentUser() {
	userID, err := StringToInt(c.UserID)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
		return
	}

	user, err := models.FindUserById(userID)
	if err != nil {
		c.renderError(err, http.StatusNotFound)
	}

	c.CurrentUser = user
}

func (c *baseAPIController) validateClientCredential() {
	clientID := c.GetString("client_id")
	clientSecret := c.GetString("client_secret")
	err := errors.New("Client authentication failed")
	client, err := oauth.ClientStore.GetByID(context.TODO(), clientID)
	if err != nil {
		c.renderError(err, http.StatusUnauthorized)
		return
	}

	if client.GetSecret() != clientSecret {
		c.renderError(err, http.StatusUnauthorized)
		return
	}
}

func (c *baseAPIController) validateTokenCredential() {
	tokenInfo, err := oauth.OauthServer.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
		return
	}

	c.UserID = tokenInfo.GetUserID()
}

func (c *baseAPIController) serveJSON(data interface{}) {
	response, err := jsonapi.Marshal(data)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	c.Data["json"] = response
	err = c.ServeJSON()
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}
}

func (c *baseAPIController) renderError(err error, status int) {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.WriteHeader(status)
	err = jsonapi.MarshalErrors(c.Ctx.ResponseWriter, []*jsonapi.ErrorObject{{
		Detail: err.Error(),
	}})
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}

	c.StopRun()
}
