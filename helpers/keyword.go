package presenters

import (
	"google-scraper/models"
)

type KeywordSearchResult struct {
	Keyword                     string
	TopAdWordAdvertisersCount   int
	TotalAdWordAdvertisersCount int
	TotalLinksCount             int
	ResultsCount                int
	Html                string
}

func KeywordPresenter(k *models.Keyword) (*KeywordSearchResult, error) {
	searchResult, err := models.FindSearchResultByKeywordId(k.Id)
	if err != nil {
		return nil, err
	}

	keywordSearchResult := KeywordSearchResult{
		Keyword:                     k.Name,
		TopAdWordAdvertisersCount:   searchResult.TopAdWordAdvertisersCount,
		TotalAdWordAdvertisersCount: searchResult.TotalAdWordAdvertisersCount,
		TotalLinksCount:             searchResult.TotalLinksCount,
		ResultsCount:                searchResult.ResultsCount,
		Html: searchResult.Html,
	}
	return &keywordSearchResult, nil
}
