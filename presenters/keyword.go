package presenters

import (
	"encoding/json"
	"errors"

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

func InitializeKeywordPresenter(k *models.Keyword) (*KeywordSearchResult, error) {
	if k == nil {
		return nil, errors.New("keyword object can't be nil")
	}

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
