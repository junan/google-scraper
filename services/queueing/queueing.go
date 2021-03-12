package queueing

import (
	"errors"

	"google-scraper/database"
	"google-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
	"github.com/beego/beego/v2/server/web"
)

var enqueuer *work.Enqueuer

func init() {
	enqueuer = work.NewEnqueuer("google_scraper", database.GetRedisPool())
}

func AddToQueue(keyword *models.Keyword, secondsInTheFuture int64) (*work.ScheduledJob, error) {
	if keyword.Id <= 0 {
		return nil, errors.New("invalid keyword object")
	}

	crawlingJobName, err := web.AppConfig.String("crawlingJobName")
	if err != nil {
		logs.Critical("crawlingJobName is not found: ", err)
		return nil, err
	}

	job, err := enqueuer.EnqueueIn(crawlingJobName, secondsInTheFuture, work.Q{"keywordId": keyword.Id})

	if err != nil {
		return nil, err
	}

	return job, nil
}
