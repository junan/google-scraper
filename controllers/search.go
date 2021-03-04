package controllers

import (
	"fmt"
	"net/http"

	"google-scraper/forms"

	"github.com/beego/beego/v2/server/web"
)

type Search struct {
	baseController
}

func (c *Search) Create() {
	file, _, err := c.GetFile("file")

	redirectPath := "/"
	flash := web.NewFlash()

	_, err = forms.SearchProcess(file)
	if err != nil {
		flash.Error(fmt.Sprint(err))
	} else {
		flash.Success("You csv file has been uploaded")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

