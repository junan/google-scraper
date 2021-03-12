package main

import (
	"os"
	"os/signal"
	"syscall"

	"google-scraper/database"
	_ "google-scraper/initializers"
	"google-scraper/worker/jobs"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocraft/work"
	"github.com/beego/beego/v2/server/web"
)

var crawlingJobName string
var enqueuerName string

func init() {
	var err error
	crawlingJobName, err = web.AppConfig.String("crawlingJobName")
	if err != nil {
		logs.Critical("crawlingJobName is not found: ", err)
	}

	enqueuerName, err = web.AppConfig.String("redisEnqueuerName")
	if err != nil {
		logs.Critical("redisEnqueuerName is not found: ", err)
	}
}

func main() {
	pool := work.NewWorkerPool(jobs.Context{}, 5, enqueuerName, database.GetRedisPool())
	pool.JobWithOptions(crawlingJobName, work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).PerformCrawling)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	// Stop the pool
	pool.Stop()
}
