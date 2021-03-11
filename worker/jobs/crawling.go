package jobs

import (
	. "google-scraper/models"
	. "google-scraper/services/crawler"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

type Context struct{}

// Number of retry
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

	logs.Info("Successfully crawled data for %v keyword", keyword.Name)

	return nil
}
