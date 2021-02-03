package fabricators

import (
	. "google-scraper/helpers"
	. "google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func FabricateUser(name string, email string, password string) {
	o := orm.NewOrm()
	hashedPassword := HashPassword(password)
	user := User{Name: name, Email: email, HashedPassword: hashedPassword}

	_, err := o.Insert(&user)
	if err != nil {
		logs.Error("User creation  failed: ", err)
	}
}
