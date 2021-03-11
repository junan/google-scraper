package keyword

import (
	"github.com/gocraft/work"

	"google-scraper/database"
	"google-scraper/models"
)

var enqueuer *work.Enqueuer

func init() {
	if enqueuer == nil {
		enqueuer = work.NewEnqueuer("google_scraper", database.GetRedisPool())
	}
}

func StartJob(keyword *models.Keyword, secondsInTheFuture int64) error {
	_, err := enqueuer.EnqueueIn("crawling_job", secondsInTheFuture, work.Q{"keywordId": keyword.Id})

	if err != nil {
		return err
	}

	return nil
}
