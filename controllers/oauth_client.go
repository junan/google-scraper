package controllers

import (
	"context"
	"fmt"
	"net/http"

	"google-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type OauthClient struct {
	baseController
}

func (c *OauthClient) New() {
	web.ReadFromRequest(&c.Controller)

	c.setClient()

	c.TplName = "oauth_client/new.html"
}

func (c *OauthClient) Create() {
	web.ReadFromRequest(&c.Controller)
	redirectPath := "/oauth-client"
	flash := web.NewFlash()
	domain := c.Ctx.Request.Host
	oauthClient, err := oauth.GenerateOauthClient(c.CurrentUser.Id, domain)
	if err != nil {
		flash.Error(err.Error())
	} else {
		redirectPath = fmt.Sprintf("/oauth-client?client_id=%s", oauthClient.ID)
		flash.Success("Your oauth client has been created successfully")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

func (c *OauthClient) setClient() {
	clientId := c.GetString("client_id")
	if clientId == "" {
		return
	}

	oauthClient, err := oauth.ClientStore.GetByID(context.TODO(), clientId)
	if err != nil {
		logs.Error("Getting oauth client failed: ", err)
	} else {
		c.Data["OauthClient"] = oauthClient
		c.Data["ClientID"] = oauthClient.GetID()
		c.Data["Secret"] = oauthClient.GetSecret()
	}
}
