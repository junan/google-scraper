package controllers

import (
	"errors"
	"net/http"

	. "google-scraper/helpers"
	. "google-scraper/models"
	"google-scraper/presenters"

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
		flash := web.NewFlash()
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}

	keywordPresenter, err := presenters.InitializeKeywordPresenter(keyword)
	if err != nil {
		logs.Error("Initializing presenter failed: ", err)
	}

	c.TplName = "keyword/show.html"
	c.Data["KeywordPresenter"] = keywordPresenter
}

func (c *KeywordController) RenderHtml() {
	web.ReadFromRequest(&c.Controller)
	keyword, err := c.findKeyword()
	if err != nil {
		flash := web.NewFlash()
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Ctx.Redirect(http.StatusFound, "/")
		return
	}

	searchResult, err := FindSearchResultByKeywordId(keyword.Id)
	if err != nil {
		logs.Error("Finding search result failed: ", err)
	}

	c.TplName = "keyword/search_result.html"
	c.Data["SearchResult"] = searchResult
}

func (c *KeywordController) findKeyword() (*Keyword, error) {
	keywordId := c.Ctx.Input.Param(":id")
	Id, err := StringToInt(keywordId)
	if err != nil {
		logs.Error("Converting String to Int failed: ", err)

		// The error message will show to the users
		return nil, errors.New("Something went wrong. Please try again.")
	}

	return FindKeywordBy(Id, c.CurrentUser)
}
