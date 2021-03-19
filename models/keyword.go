package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Keyword struct {
	Base

	Name            string
	SearchCompleted bool  `default:"false"`
	User            *User `orm:"rel(fk)"`
}

func CreateKeyword(k *Keyword) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(k)
}

func UpdateKeyword(k *Keyword) (*Keyword, error) {
	orm := orm.NewOrm()
	_, err := orm.Update(k, "SearchCompleted")
	if err != nil {
		return nil, err
	}

	return k, nil
}

func GetKeywords(u *User) orm.QuerySeter {
	orm := orm.NewOrm()
	return orm.QueryTable(Keyword{}).Filter("user_id", u.Id)
}

func GetPaginatedKeywords(u *User, offset int, sizePerPage int) ([]*Keyword, error) {
	userKeywords := []*Keyword{}

	_, err := GetKeywords(u).Limit(sizePerPage, offset).OrderBy("-id").All(&userKeywords)
	if err != nil {
		logs.Error("Keywords pagination failed: ", err)
	}

	return userKeywords, nil
}

func FindKeywordById(id int64) (keyword *Keyword, err error) {
	orm := orm.NewOrm()
	keyword = &Keyword{
		Base: Base{Id: id},
	}

	err = orm.Read(keyword)
	if err != nil {
		return nil, err
	}

	return keyword, nil
}

func (k *Keyword) CreatedByUser(u *User) bool {
	userKeywords := GetKeywords(u)
	var keyword Keyword
	err := userKeywords.Filter("id", k.Id).One(&keyword)
	if err != nil {
		return false
	}

	return true
}

// Beego by default creates the table name as singular, it will make it plural
func (k *Keyword) TableName() string {
	return "keywords"
}
