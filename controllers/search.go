package controllers

import (
	"net/http"

	"google-scraper/forms"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
)

type Search struct {
	baseController
}

func (c *Search) Create() {
	redirectPath := "/"
	flash := web.NewFlash()
	file, header, err := c.GetFile("file")

	if err != nil {
		logs.Error("Getting file failed: ", err)
	}
	err = forms.PerformSearch(file, header, c.CurrentUser)

	if err != nil {
		flash.Error(err.Error())
	} else {
		flash.Success("Your csv file has been uploaded successfully")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}
