package serializers

import (
	"time"

	"github.com/tidwall/gjson"

	"github.com/beego/beego/v2/core/utils/pagination"

	"google-scraper/models"
)

type KeywordList struct {
	Keywords []*models.Keyword

	Paginator *pagination.Paginator
}

type KeywordListResponse struct {
	Id              int64  `jsonapi:"primary,keyword"`
	Name            string `jsonapi:"attr,name"`
	SearchCompleted bool   `jsonapi:"attr,search_completed"`
	CreatedAt       string `jsonapi:"attr,created_at"`
}

func (serializer *KeywordList) Data() []*KeywordListResponse {
	var data []*KeywordListResponse

	for _, keyword := range serializer.Keywords {
		data = append(data, createKeywordResponse(keyword))
	}

	return data
}

func createKeywordResponse(keyword *models.Keyword) *KeywordListResponse {
	return &KeywordListResponse{
		Id:        keyword.Id,
		Keyword:   keyword.Keyword,
		Status:    string(keyword.Status),
		CreatedAt: keyword.CreatedAt.Format(time.RFC3339),
	}
}