package jobs

import (
	"github.com/gocraft/work"

	. "google-scraper/models"
	. "google-scraper/services/crawler"

	"github.com/beego/beego/v2/core/logs"
)

type Context struct{}

const MaxFails = 3

func (c *Context) PerformCrawling(job *work.Job) error {
	keywordId := job.ArgInt64("keywordId")
	keyword, err := FindKeywordById(keywordId)
	if err != nil {
		logs.Error("Finding keyword failed: ", err)
		return err
	}

	_, err = Crawl(keyword)
	if err != nil {
		logs.Error("Crawling failed: ", err)
		return err
	}

	return nil
}
