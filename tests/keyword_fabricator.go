package tests

import (
	. "google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func FabricateKeyword(name string, searchCompleted bool, user *User) Keyword {
	o := orm.NewOrm()
	keyword := Keyword{Name: name, User: user, SearchCompleted: searchCompleted}

	_, err := o.Insert(&keyword)
	if err != nil {
		logs.Error("Creating keyword failed: ", err)
	}

	return keyword
}
