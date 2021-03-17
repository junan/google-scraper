package enqueueing

import (
	"fmt"
	"errors"

	"google-scraper/database"
	"google-scraper/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gocraft/work"
)

var enqueuer *work.Enqueuer

func init() {
	enqueuer = work.NewEnqueuer("google_scraper", database.GetRedisPool())
}

func EnqueueKeywordJob(keyword *models.Keyword, throttleMultiplier int64) (*work.ScheduledJob, error) {
	if keyword == nil {
		return nil,  errors.New("keyword object can't be nil")
	}

	if keyword.Id <= 0 {
		return nil,  fmt.Errorf("invalid keyword object: %+v", keyword)
	}

	crawlingJobName, err := web.AppConfig.String("crawlingJobName")
	if err != nil {
		logs.Critical("crawlingJobName is not found: ", err)
		return nil, err
	}

	delayTimeInSeconds := getDelayTimeInSeconds(throttleMultiplier)
	job, err := enqueuer.EnqueueIn(crawlingJobName, delayTimeInSeconds, work.Q{"keywordId": keyword.Id})

	if err != nil {
		return nil, err
	}

	return job, nil
}

func getDelayTimeInSeconds(throttleMultiplier int64) int64 {
	// Each job will be run two seconds later than the previous job, jobs will be enqueued immediately
	// But will be run based on this seconds value in the future
	return throttleMultiplier + 2
}
