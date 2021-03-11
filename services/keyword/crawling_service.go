package keyword

import (
	"github.com/gocraft/work"

	"google-scraper/database"
	"google-scraper/models"
)

var enqueuer *work.Enqueuer

func StartJob(keyword *models.Keyword) error {

	setUpEnqueuer()
	_, err := enqueuer.Enqueue("crawling_job", work.Q{"keywordId": keyword.Id})

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
