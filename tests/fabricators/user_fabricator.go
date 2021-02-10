package fabricators

import (
	. "google-scraper/helpers"
	. "google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func FabricateUser(name string, email string, password string) (User, error) {
	o := orm.NewOrm()
	hashedPassword, err := HashPassword(password)
	if err != nil {
		logs.Error("Password hashing failed: ", err)
	}
	user := User{Name: name, Email: email, HashedPassword: hashedPassword}

	_, err = o.Insert(&user)
	if err != nil {
		logs.Error("User creation  failed: ", err)
	}

	return user, err
}
