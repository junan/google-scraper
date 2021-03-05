package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

type Keyword struct {
	Base

	Name string
	SearchCompleted bool  `default:"false"`
	User *User `orm:"rel(fk)"`
}

func CreateKeyword(s *Keyword) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(s)
}

// Beego by default creates the table name as singular, it will make it plural
func (u *Keyword) TableName() string {
	return "keywords"
}
