package models

import (
	"errors"

	"github.com/beego/beego/v2/adapter/orm"
)

type Keyword struct {
	Base

	Name            string
	SearchCompleted bool  `default:"false"`
	User            *User `orm:"rel(fk)"`
}

func CreateKeyword(s *Keyword) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(s)
}


func UpdateKeyword(s *Keyword) (int64, error) {
	orm := orm.NewOrm()
	return orm.Update(s, "SearchCompleted")
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

func FetchKeywords(user *User) ([]*Keyword, error) {
	if user == nil {
		return nil, errors.New("User is blank")
	}

	var keywords []*Keyword

	_, err := userKeywordsQuerySeter(user).OrderBy("-id").All(&keywords)

	return keywords, err
}

// Beego by default creates the table name as singular, it will make it plural
func (u *Keyword) TableName() string {
	return "keywords"
}

func userKeywordsQuerySeter(user *User) orm.QuerySeter {
	o := orm.NewOrm()

	return o.QueryTable(Keyword{}).Filter("user_id", user.Id)
}
