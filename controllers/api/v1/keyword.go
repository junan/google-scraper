package apiv1

import (
	"net/http"

	. "google-scraper/helpers"
	"google-scraper/models"
	"google-scraper/presenters"
	"google-scraper/serializers"

	"github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/adapter/utils/pagination"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
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
		c.renderError(err, http.StatusInternalServerError)
	}

	paginator := pagination.SetPaginator((*context.Context)(c.Ctx), sizePerPage, keywordsCount)
	paginatedKeywords, err := models.GetPaginatedKeywords(keywords, paginator.Offset(), sizePerPage)

	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	keywordsSerializer := serializers.KeywordList{
		Keywords:  paginatedKeywords,
		Paginator: paginator,
	}

	c.serveListJSON(keywordsSerializer.Data(), keywordsSerializer.Meta(), keywordsSerializer.Links())
}

func (c *Keyword) Show() {
	keyword, err := c.findKeyword()
	if err != nil {
		c.renderError(err, http.StatusNotFound)
	}

	keywordResult, err := presenters.InitializeKeywordPresenter(keyword)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	response := serializers.GetKeywordResponse(*keywordResult)
	c.serveJSON(response)
}

func (c *Keyword) findKeyword() (*models.Keyword, error) {
	keywordId := c.Ctx.Input.Param(":id")
	Id, err := StringToInt(keywordId)
	if err != nil {
		c.renderError(err, http.StatusInternalServerError)
	}

	return models.FindKeywordBy(Id, c.CurrentUser)
}
