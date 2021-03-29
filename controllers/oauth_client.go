package controllers

import (
	"fmt"
	"context"
	"net/http"

	"github.com/beego/beego/v2/server/web"

	"google-scraper/models"
	"google-scraper/services/oauth"

	"github.com/beego/beego/v2/core/logs"
)

type OauthClient struct {
	baseController
}

func (c *OauthClient) New() {
	web.ReadFromRequest(&c.Controller)
	flash := web.NewFlash()

	clientId := c.GetString("client_id")

	if clientId != "" {
		oauthClient, err := oauth.ClientStore.GetByID(context.TODO(), clientId)
		if err != nil {
			flash.Error(err.Error())
		} else {
			c.Data["OauthClient"] = oauthClient
		}
	}

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
		logs.Error("Retrieving keyword count failed: ", err, oauthClient)
	} else {
		redirectPath = fmt.Sprintf("/oauth-client?client_id=%s", oauthClient.ID)
		flash.Success("Your oauth client has been created successfully")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

func (c *Dashboard) OauthClient(paginatedKeywords []*models.Keyword, keyword string) {
	c.TplName = "dashboard/new.html"
	c.Data["Keywords"] = paginatedKeywords
	c.Data["SearchKeyword"] = keyword
}
