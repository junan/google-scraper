package serializers

import (
	. "google-scraper/presenters"
)

type KeywordResponse struct {
	Id                          int64    `jsonapi:"primary,keyword_detail"`
	Name                        string   `jsonapi:"attr,name"`
	CreatedAt                   string   `jsonapi:"attr,created_at"`
	TopAdWordAdvertisersCount   int      `jsonapi:"attr,top_ad_word_advertisers_count"`
	TotalAdWordAdvertisersCount int      `jsonapi:"attr,total_ad_word_advertisers_count"`
	TotalLinksCount             int      `jsonapi:"attr,total_links_count"`
	ResultsCount                int      `jsonapi:"attr,results_count"`
	TopAdWordAdvertisersUrls    []string `jsonapi:"attr,top_ad_word_advertisers_urls"`
	ResultsUrls                 []string `jsonapi:"attr,results_url"`
	Html                        string   `jsonapi:"attr,html"`
}

func GetKeywordResponse(keywordResult KeywordSearchResult) *KeywordResponse {
	return &KeywordResponse{
		Id:                          keywordResult.KeywordId,
		Name:                        keywordResult.Keyword,
		CreatedAt:                   keywordResult.CreatedAt,
		TopAdWordAdvertisersCount:   keywordResult.TopAdWordAdvertisersCount,
		TotalAdWordAdvertisersCount: keywordResult.TotalAdWordAdvertisersCount,
		TotalLinksCount:             keywordResult.TotalLinksCount,
		ResultsCount:                keywordResult.ResultsCount,
		Html:                        keywordResult.Html,
		TopAdWordAdvertisersUrls:    keywordResult.TopAdWordAdvertisersUrls,
		ResultsUrls:                 keywordResult.ResultsUrls,
	}
}
