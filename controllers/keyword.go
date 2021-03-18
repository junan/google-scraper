package controllers

import (
	. "google-scraper/helpers"
	. "google-scraper/models"

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

	c.setAttributes()
}

func (c *KeywordController) setAttributes() {
	c.TplName = "keyword/show.html"
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
