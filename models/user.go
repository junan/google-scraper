package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

type User struct {
	Base

	Name           string
	Email          string `orm:"unique"`
	HashedPassword string
}

func (u *User) Keywords() orm.QuerySeter {
	orm := orm.NewOrm()
	return orm.QueryTable(Keyword{}).Filter("user_id", u.Id)
}

func (u *User) PaginatedKeywords(keywords orm.QuerySeter, offset int, paginatesPer int) ([]*Keyword, error) {
	userKeywords := []*Keyword{}

	_, err := keywords.Limit(paginatesPer, offset).OrderBy("-id").All(&userKeywords)
	if err != nil {
		logs.Error("Keywords pagination failed: ", err)
	}

	return userKeywords, nil
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
