package models

import (
	"errors"

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

func GetKeywords(u *User, keyword string) orm.QuerySeter {
	orm := orm.NewOrm()
	result :=  orm.QueryTable(Keyword{}).Filter("user_id", u.Id)

	if len(keyword) > 0 {
		result = result.Filter("name__icontains", keyword)
	}

	return result
}

func GetPaginatedKeywords(keywords orm.QuerySeter, offset int, sizePerPage int) ([]*Keyword, error) {
	userKeywords := []*Keyword{}

	_, err := keywords.Limit(sizePerPage, offset).OrderBy("-id").All(&userKeywords)
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

func FindKeywordBy(Id int64, u *User) (*Keyword, error) {
	var keyword Keyword
	userKeywords := GetKeywords(u, "")
	err := userKeywords.Filter("id", Id).One(&keyword)
	if err != nil {
		return nil, errors.New("Keyword not found.")
	}

	return &keyword, nil
}

// Beego by default creates the table name as singular, it will make it plural
func (k *Keyword) TableName() string {
	return "keywords"
}
