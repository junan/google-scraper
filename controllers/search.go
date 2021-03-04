package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"google-scraper/forms"

	"github.com/beego/beego/v2/server/web"
)

type Search struct {
	baseController
}

func (c *Search) Create() {

	//fileExt := path.Ext(head.Filename)
	//if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg"{
	//	beego.Info ("The uploaded image format is incorrect, please add it again!")
	//	this.TplName = "add.html"
	//	return
	//}

	records, err := c.readData()

	fmt.Println(records)

	//fmt.Println(content)
	searchForm := forms.SearchForm{}
	flash := web.NewFlash()
	redirectPath := "/"

	err = c.ParseForm(&searchForm)
	if err != nil {
		flash.Error(err.Error())
	}

	_, err = searchForm.Save(c.CurrentUser)
	if err != nil {
		flash.Error(fmt.Sprint(err))
	} else {
		flash.Success("You csv file has been uploaded")
	}

	flash.Store(&c.Controller)
	c.Ctx.Redirect(http.StatusFound, redirectPath)
}

func (c *Search) readData() ([][]string, error) {
	file, _, err := c.GetFile("file")

	r := csv.NewReader(file)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
