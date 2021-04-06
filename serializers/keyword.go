package serializers

import (
	. "google-scraper/helpers"
	"google-scraper/models"

	"github.com/beego/beego/v2/adapter/utils/pagination"
	"github.com/google/jsonapi"
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

	for _, k := range serializer.Keywords {
		data = append(data, createKeywordResponse(k))
	}

	return data
}

func (serializer *KeywordList) Meta() (meta *jsonapi.Meta) {
	return &jsonapi.Meta{
		"page":    serializer.Paginator.Page(),
		"pages":   serializer.Paginator.PageNums(),
		"records": serializer.Paginator.Nums(),
	}
}

func (serializer *KeywordList) Links() (links *jsonapi.Links) {
	return &jsonapi.Links{
		"self":  serializer.Paginator.PageLink(serializer.Paginator.Page()),
		"first": serializer.Paginator.PageLinkFirst(),
		"prev":  serializer.Paginator.PageLinkPrev(),
		"next":  serializer.Paginator.PageLinkNext(),
		"last":  serializer.Paginator.PageLinkLast(),
	}
}

func createKeywordResponse(keyword *models.Keyword) *KeywordListResponse {
	return &KeywordListResponse{
		Id:        keyword.Id,
		Name:   keyword.Name,
		SearchCompleted: keyword.SearchCompleted,
		CreatedAt: DisplayFormattedCreatedDate(*keyword),
	}
}
