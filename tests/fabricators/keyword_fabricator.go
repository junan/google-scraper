package fabricators

import (
	. "google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func FabricateKeyword(name string, user *User) Keyword {
	o := orm.NewOrm()
	keyword := Keyword{Name: name, User: user}

	_, err := o.Insert(&keyword)
	if err != nil {
		logs.Error("keyword creation  failed: ", err)
	}

	return keyword
}
