package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

type User struct {
	Base

	Name              string
	Email             string `orm:"unique"`
	HashedPassword    string
}

func CreateUser(u *User)(id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(u)
}

// Beego by default creates the table name as singular, it will make it plural
func (u *User) TableName() string {
	return "users"
}
