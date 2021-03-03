package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

type Search struct {
	Base

	Keyword string
	User *User `orm:"rel(fk)"`
}

func CreateResult(s *Search) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(s)
}

// Beego by default creates the table name as singular, it will make it plural
func (u *Search) TableName() string {
	return "searches"
}
