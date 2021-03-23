package controllers

import (
	"google-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/pagination"
)

var sizePerPage int

type Dashboard struct {
	baseController
}

func init() {
	var err error
	sizePerPage, err = web.AppConfig.Int("sizePerPage")
	if err != nil {
		logs.Error("Retrieving sizePerPage failed: ", err)
	}
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)

	searchedKeyword := c.GetString("keyword")
	keywords := models.GetKeywords(c.CurrentUser, searchedKeyword)

	keywordsCount, err := keywords.Count()
	if err != nil {
		logs.Error("Retrieving keyword count failed: ", err)
	}

	paginator := pagination.SetPaginator(c.Ctx, sizePerPage, keywordsCount)
	paginatedKeywords, err := models.GetPaginatedKeywords(keywords, paginator.Offset(), sizePerPage)
	if err != nil {
		logs.Error("Retrieving keywords failed: ", err)
	}

	c.setAttributes(paginatedKeywords, searchedKeyword)
}

func (c *Dashboard) setAttributes(paginatedKeywords []*models.Keyword, searchedKeyword string) {
	c.TplName = "dashboard/new.html"
	c.Data["Keywords"] = paginatedKeywords
	c.Data["SearchKeyword"] = searchedKeyword
}
