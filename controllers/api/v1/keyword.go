package apiv1

import (
	"google-scraper/models"
	"google-scraper/serializers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/adapter/utils/pagination"
)

var sizePerPage int

type Keyword struct {
	baseAPIController
}

func init() {
	var err error
	sizePerPage, err = web.AppConfig.Int("sizePerPage")
	if err != nil {
		logs.Error("Retrieving sizePerPage failed: ", err)
	}
}

func (c *Keyword) Index() {
	keyword := c.GetString("keyword")
	keywords := models.GetQuerySeterKeywords(c.CurrentUser, keyword)

	keywordsCount, err := keywords.Count()
	if err != nil {
		logs.Error("Retrieving keyword count failed: ", err)
	}

	paginator := pagination.SetPaginator((*context.Context)(c.Ctx), sizePerPage, keywordsCount)
	paginatedKeywords, err := models.GetPaginatedKeywords(keywords, paginator.Offset(), sizePerPage)

	if err != nil {
		logs.Error("Retrieving keywords failed: ", err)
	}

	keywordsSerializer := serializers.KeywordList{
		Keywords:  paginatedKeywords,
		Paginator: paginator,
	}

	c.serveListJSON(keywordsSerializer.Data(), keywordsSerializer.Meta(), keywordsSerializer.Links())
}
