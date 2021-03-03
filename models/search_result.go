package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

type SearchResult struct {
	Base

	TopAdWordAdvertisersCount   int `default:"0"`
	TopAdWordAdvertisersUrls    string `orm:"type(json);"`
	TotalAdWordAdvertisersCount int `default:"0"`
	ResultsCount                int `default:"0"`
	ResultsUrls                 string `orm:"type(json);"`
	TotalLinksCount             int `default:"0"`
	Html                        string
}

func CreateSearchResult(s *SearchResult) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(s)
}
//
//func (u *User) IsExistingUser() bool {
//	o := orm.NewOrm()
//	err := o.Read(u, "Email")
//	return err == nil
//}
//
//func FindUserByEmail(email string) (user *User, err error) {
//	orm := orm.NewOrm()
//	user = &User{Email: email}
//
//	err = orm.Read(user, "Email")
//	if err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}
//
//func FindUserById(id int64) (user *User, err error) {
//	orm := orm.NewOrm()
//	user = &User{
//		Base: Base{Id: id},
//	}
//
//	err = orm.Read(user)
//	if err != nil {
//		return nil, err
//	}
//
//	return user, nil
//}

// Beego by default creates the table name as singular, it will make it plural
func (u *SearchResult) TableName() string {
	return "search_results"
}
