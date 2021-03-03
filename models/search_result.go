package models

import (
	"github.com/beego/beego/v2/adapter/orm"
)

type SearchResult struct {
	Base

	TopAdWordAdvertisersCount   int    `default:"0"`
	TopAdWordAdvertisersUrls    string `orm:"type(json);null"`
	TotalAdWordAdvertisersCount int    `default:"0"`
	ResultsCount                int    `default:"0"`
	ResultsUrls                 string `orm:"type(json);null"`
	TotalLinksCount             int    `default:"0"`
	Html                        string `orm:"type(text);"`
	Search *Search `orm:"rel(fk)"`
}

func CreateSearchResult(s *SearchResult) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(s)
}

// Beego by default creates the table name as singular, it will make it plural
func (u *SearchResult) TableName() string {
	return "search_results"
}
