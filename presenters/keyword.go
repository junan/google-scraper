package presenters

import (
	"encoding/json"

	"google-scraper/models"
)

type KeywordSearchResult struct {
	KeywordId                   int64
	Keyword                     string
	TopAdWordAdvertisersCount   int
	TotalAdWordAdvertisersCount int
	TotalLinksCount             int
	ResultsCount                int
	Html                        string
	TopAdWordAdvertisersUrls    []string
	ResultsUrls                 []string
}

func KeywordPresenter(k *models.Keyword) (*KeywordSearchResult, error) {
	searchResult, err := models.FindSearchResultByKeywordId(k.Id)
	if err != nil {
		return nil, err
	}

	topAdWordAdvertisersUrls, err := unmarshalUrls(searchResult.TopAdWordAdvertisersUrls)
	if err != nil {
		return nil, err
	}

	resultsUrls, err := unmarshalUrls(searchResult.ResultsUrls)
	if err != nil {
		return nil, err
	}

	keywordSearchResult := KeywordSearchResult{
		KeywordId:                   k.Id,
		Keyword:                     k.Name,
		TopAdWordAdvertisersCount:   searchResult.TopAdWordAdvertisersCount,
		TotalAdWordAdvertisersCount: searchResult.TotalAdWordAdvertisersCount,
		TotalLinksCount:             searchResult.TotalLinksCount,
		ResultsCount:                searchResult.ResultsCount,
		Html:                        searchResult.Html,
		TopAdWordAdvertisersUrls:    topAdWordAdvertisersUrls,
		ResultsUrls:                 resultsUrls,
	}

	return &keywordSearchResult, nil
}

func unmarshalUrls(urls string) ([]string, error) {
	var result []string
	err := json.Unmarshal([]byte(urls), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
