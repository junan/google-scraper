package controllers

import (
	"net/http"

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
	keyword, err := c.findKeyword()
	if err != nil {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}

	result := keyword.IsBelongTo(c.CurrentUser)
	if !result {
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}

	keywordPresenter, err := presenter.KeywordPresenter(keyword)
	if err != nil {
		logs.Error("Initializing presenter failed: ", err)
	}

	c.setAttributes(keywordPresenter)
}

func (c *KeywordController) findKeyword() (*Keyword, error) {
	keywordId := c.Ctx.Input.Param(":id")
	Id, err := StringToInt(keywordId)
	if err != nil {
		logs.Error("Converting String to Int failed: ", err)
		return nil, err
	}

	return FindKeywordById(Id)
}

func (c *KeywordController) setAttributes(keywordPresenter *presenter.KeywordSearchResult) {
	c.TplName = "keyword/show.html"
	c.Data["KeywordPresenter"] = keywordPresenter
}
