package queueing

import (
	"github.com/gocraft/work"

	"google-scraper/database"
	"google-scraper/models"
)

var enqueuer *work.Enqueuer

func init() {
	enqueuer = work.NewEnqueuer("google_scraper", database.GetRedisPool())
}

func AddToQueue(keyword *models.Keyword, secondsInTheFuture int64) error {
	_, err := enqueuer.EnqueueIn("crawling_job", secondsInTheFuture, work.Q{"keywordId": keyword.Id})

	if err != nil {
		return err
	}

	return nil
}
