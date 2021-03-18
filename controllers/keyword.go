package controllers

import (
	. "google-scraper/helpers"
	. "google-scraper/models"
	presenter "google-scraper/presenters"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type KeywordController struct {
	baseController
}

func (c *KeywordController) Show() {
	web.ReadFromRequest(&c.Controller)
	keyword, err := c.getKeyword()
	if keyword != nil {
		logs.Error("Converting String to Int failed: ", err)
	}

	keywordPresenter, err := presenter.KeywordPresenter(keyword)

	c.setAttributes(keywordPresenter)
}

func (c *KeywordController) getKeyword() (*Keyword, error) {
	keywordId := c.Ctx.Input.Param(":id")
	Id, err := StringToInt(keywordId)
	if err != nil {
		logs.Error("Converting String to Int failed: ", err)
		return nil, err
	}

	return FindKeywordById(Id)
}

func (c *KeywordController) setAttributes(ksr *presenter.KeywordSearchResult) {
	c.TplName = "keyword/show.html"
	c.Data["Keywords"] = ksr
}

