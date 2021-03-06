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
	var user User

	err := o.QueryTable(User{}).Filter("email__iexact", u.Email).One(&user)

	return err == nil
}

func FindUserByEmail(email string) (user *User, err error) {
	orm := orm.NewOrm()
	user = &User{Email: email}

	err = orm.Read(user, "Email")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserById(id int64) (user *User, err error) {
	orm := orm.NewOrm()
	user = &User{
		Base: Base{Id: id},
	}

	err = orm.Read(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Beego by default creates the table name as singular, it will make it plural
func (u *User) TableName() string {
	return "users"
}
