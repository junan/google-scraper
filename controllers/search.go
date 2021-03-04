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
	file, header, _ := c.GetFile("file")
	redirectPath := "/"
	flash := web.NewFlash()

	errs := forms.SearchProcess(file, header)

	if len(errs) > 0 {
		flash.Error(fmt.Sprint(errs))
	} else {
		flash.Success("You csv file has been uploaded")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

