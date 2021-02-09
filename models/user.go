package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

type User struct {
	Base

	Name           string
	Email          string `orm:"unique"`
	HashedPassword string
}

func CreateUser(u *User) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(u)
}

func (u *User) IsExistingUser() bool {
	o := orm.NewOrm()
	err := o.Read(u, "Email")
	return err == nil
}

// Beego by default creates the table name as singular, it will make it plural
func (u *User) TableName() string {
	return "users"
}
