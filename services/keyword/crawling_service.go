package keyword

import (
	"github.com/gocraft/work"

	"google-scraper/database"
	"google-scraper/models"
)

var enqueuer *work.Enqueuer

func StartJob(keyword *models.Keyword, secondsInTheFuture int64) error {
	setUpEnqueuer()
	_, err := enqueuer.EnqueueIn("crawling_job", secondsInTheFuture, work.Q{"keywordId": keyword.Id})

	if err != nil {
		return err
	}

	return nil
}

func setUpEnqueuer() {
	if enqueuer == nil {
		enqueuer = work.NewEnqueuer("google_scraper", database.GetRedisPool())
	}
}
