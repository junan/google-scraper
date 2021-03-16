package controllers

import (
	"google-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/pagination"
)

var paginatesPer int

type Dashboard struct {
	baseController
}

func init() {
	var err error
	paginatesPer, err = web.AppConfig.Int("paginatesPer")
	if err != nil {
		logs.Error("Retrieving paginatesPer failed: ", err)
	}
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)
	keywords := c.CurrentUser.Keywords()
	keywordsCount, err := keywords.Count()
	if err != nil {
		logs.Error("Retrieving keyword count failed: ", err)
	}

	paginator := pagination.SetPaginator(c.Ctx, paginatesPer, keywordsCount)
	paginatedKeywords, err := c.CurrentUser.PaginatedKeywords(keywords, paginator.Offset(), paginatesPer)
	if err != nil {
		logs.Error("Retrieving keywords failed: ", err)
	}

	c.setAttributes(paginatedKeywords)
}

func (c *Dashboard) setAttributes(paginatedKeywords []*models.Keyword) {
	c.TplName = "dashboard/new.html"
	c.Data["Keywords"] = paginatedKeywords
}
