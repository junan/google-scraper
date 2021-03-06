package apiv1

import (
	"context"
	"errors"
	"net/http"

	. "google-scraper/helpers"
	"google-scraper/models"
	"google-scraper/services/oauth"

	"github.com/beego/beego/v2/server/web"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/jsonapi"
)

type baseAPIController struct {
	web.Controller

	CurrentUser *models.User
	authPolicy  AuthPolicy
	actionName  string
	TokenInfo   oauth2.TokenInfo
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
	c.setActionName()

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}

	c.validateAuthorization()
	c.setCurrentUser()
}

func (c *baseAPIController) setDefaultAuthPolicy() {
	// By default it will validate API token and will not validate client credential for all routes.
	// Override this default policy in the `NestPrepare()` function when necessary
	c.authPolicy = AuthPolicy{requireTokenValidation: true, requireClientCredentialValidation: false}
}

func (c *baseAPIController) validateAuthorization() {
	err := errors.New("Client authentication failed")
	if c.authPolicy.requireClientCredentialValidation {
		result := c.validClientCredential()
		if !result {
			c.renderError(err, http.StatusUnauthorized)
		}
	}

	if c.authPolicy.requireTokenValidation {
		result := c.validToken()
		if !result {
			c.renderError(err, http.StatusUnauthorized)
		}
	}
}

func (c *baseAPIController) setCurrentUser() {
	if c.TokenInfo == nil {
		return
	}

	userID, err := StringToInt(c.TokenInfo.GetUserID())
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	user, err := models.FindUserById(userID)
	if err != nil {
		c.renderError(err, http.StatusNotFound)
	}

	c.CurrentUser = user
}

func (c *baseAPIController) setActionName() {
	_, actionName := c.GetControllerAndAction()
	c.actionName = actionName
}

func (c *baseAPIController) validClientCredential() bool {
	clientID := c.GetString("client_id")
	clientSecret := c.GetString("client_secret")

	if clientID == "" {
		return false
	}

	client, err := oauth.ClientStore.GetByID(context.TODO(), clientID)
	if err != nil {
		return false
	}

	if client.GetSecret() != clientSecret {
		return false
	}

	return true
}

func (c *baseAPIController) validToken() bool {
	tokenInfo, err := oauth.OauthServer.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		return false
	}
	c.TokenInfo = tokenInfo

	return true
}

func (c *baseAPIController) serveJSON(data interface{}) {
	payload, err := jsonapi.Marshal(data)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	c.renderPayload(payload)
}

func (c *baseAPIController) serveListJSON(data interface{}, meta *jsonapi.Meta, links *jsonapi.Links) {
	response, err := jsonapi.Marshal(data)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	payload, ok := response.(*jsonapi.ManyPayload)
	if !ok {
		c.renderError(err, http.StatusInternalServerError)
	}

	payload.Meta = meta
	payload.Links = links

	c.renderPayload(payload)
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

func (c *baseAPIController) renderPayload(payload interface{}) {
	c.Data["json"] = payload
	err := c.ServeJSON()
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}
}
