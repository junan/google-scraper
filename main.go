package main

import (
	"fmt"
	_ "google-scraper/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/lib/pq"
)

// Model Struct
type User struct {
	ID   int
	Name string `orm:"size(100)"`
}

func init() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		fmt.Println("Postgres Driver registration failed: ", err)
	}

	err = orm.RegisterDataBase("default", "postgres", "postgresql://postgres:postgres@0.0.0.0:5432/google_scraper_development?sslmode=disable")
	if err != nil {
		fmt.Println("Postgres connets failed: ", err)
	}

	// register model
	orm.RegisterModel(new(User))

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical(fmt.Sprintf("Failed to sync the database %v", err))
	}
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// // read one
	// u := User{ID: user.ID}
	// err = o.Read(&u)
	// fmt.Printf("ERR: %v\n", err)

	// // delete
	// num, err = o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	beego.Run()
}
