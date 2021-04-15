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

	//One to one relationship, each keyword(model) has one search result(model)
	Keyword *Keyword `orm:"null;rel(one);on_delete(set_null)"`
}

func CreateSearchResult(s *SearchResult) (id int64, err error) {
	orm := orm.NewOrm()
	return orm.Insert(s)
}

func FindSearchResultByKeywordId(keywordId int64) (*SearchResult, error) {
	orm := orm.NewOrm()
	var searchResult SearchResult
	err := orm.QueryTable("search_results").Filter("keyword_id", keywordId).One(&searchResult)
	if err != nil {
		return nil, err
	}

	return &searchResult, nil
}

// Beego by default creates the table name as singular, it will make it plural
func (u *SearchResult) TableName() string {
	return "search_results"
}
