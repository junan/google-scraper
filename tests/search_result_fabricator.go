package tests

import (
	. "google-scraper/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func FabricateSearchResult(k *Keyword) SearchResult {
	o := orm.NewOrm()
	searchResult := SearchResult{
		TopAdWordAdvertisersCount:   2,
		TopAdWordAdvertisersUrls:    `["http://example1.com", "http://example2.com"]`,
		TotalAdWordAdvertisersCount: 3,
		ResultsCount:                2,
		ResultsUrls:                 `["http://example1.com", "http://example2.com"]`,
		TotalLinksCount:             20,
		Html:                        "html-response",
		Keyword:                     k,
	}

	_, err := o.Insert(&searchResult)
	if err != nil {
		logs.Error("Creating SearchResult failed: ", err)
	}

	return searchResult
}
