package jobs

import (
	"github.com/gocraft/work"

	"google-scraper/models"
	. "google-scraper/services/crawler"
	"google-scraper/database"

	"github.com/beego/beego/v2/core/logs"
)

type Context struct{}

var enqueuer *work.Enqueuer

type Crawling struct {
	CsvKeywords [][]string
	UserId int64
}

func (c *Context) PerformCrawling(job *work.Job) error {
	keywords := job.ArgString("keywords")
	for _, row := range keywords {
		for _, name := range row {
			keyword, err := storeKeyword(name, crawling.UserId)
			if err == nil {
				_, err = Crawl (keyword)
				if err != nil {
					logs.Error("Crawling failed: ", err)
				}
			}
		}
	}
}

