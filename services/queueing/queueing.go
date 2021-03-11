package queueing

import (
	"google-scraper/database"
	"google-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
)

var enqueuer *work.Enqueuer

func init() {
	enqueuer = work.NewEnqueuer("google_scraper", database.GetRedisPool())
}

func AddToQueue(keyword *models.Keyword, secondsInTheFuture int64) error {
	job, err := enqueuer.EnqueueIn("crawling_job", secondsInTheFuture, work.Q{"keywordId": keyword.Id})

	if err != nil {
		return err
	}

	logs.Info("Enqueued %v keyword to the %v", keyword.Name, job.Name)

	return nil
}
