package fabricators

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	. "google-scraper/helpers"
	. "google-scraper/models"
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
