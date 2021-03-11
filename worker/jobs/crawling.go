package jobs

import (
	"github.com/gocraft/work"

	"github.com/beego/beego/v2/core/logs"

	. "google-scraper/models"
	. "google-scraper/services/crawler"
)

type Context struct{}

const MaxFails = 3

var enqueuer *work.Enqueuer

func (c *Context) PerformCrawling(job *work.Job) error {
	keywordId := job.ArgInt64("keywordId")
	keyword, err := FindKeywordById(keywordId)
	if err != nil {
		logs.Error("Finding user failed: ", err)
		return err
	}

	_, err = Crawl(keyword)
	if err != nil {
		logs.Error("Crawling failed: ", err)
		return err
	}

	return nil
}
