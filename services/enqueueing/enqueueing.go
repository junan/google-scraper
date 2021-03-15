package enqueueing

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

func DelayedEnqueue(keyword *models.Keyword, keywordIndex int64) (*work.ScheduledJob, error) {
	if keyword.Id <= 0 {
		return nil, errors.New("invalid keyword object")
	}

	crawlingJobName, err := web.AppConfig.String("crawlingJobName")
	if err != nil {
		logs.Critical("crawlingJobName is not found: ", err)
		return nil, err
	}

	delayTimeInSeconds := getDelayTimeInSeconds(keywordIndex)
	job, err := enqueuer.EnqueueIn(crawlingJobName, delayTimeInSeconds, work.Q{"keywordId": keyword.Id})

	if err != nil {
		return nil, err
	}

	return job, nil
}

func getDelayTimeInSeconds(keywordIndex int64) int64 {
	// Each job will be run two seconds later than the previous job, jobs will be enqueued immediately
	// But will be run based on this seconds value in the future
	return keywordIndex + 2
}
