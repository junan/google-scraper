package models

import (
	"github.com/beego/beego/v2/adapter/orm"
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

// Beego by default creates the table name as singular, it will make it plural
func (k *Keyword) TableName() string {
	return "keywords"
}
