package main

import (
	"fmt"
	"google-scraper/models"

	_ "google-scraper/db"
	_ "google-scraper/models"
	_ "google-scraper/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	o := orm.NewOrm()

	user := models.User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)


	beego.Run()
}
