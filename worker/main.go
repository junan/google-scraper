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
)

func init() {
}

func main() {
	logs.Error("Main worker is running: ")
	pool := work.NewWorkerPool(jobs.Context{}, 5, "google_scraper", database.GetRedisPool())

	pool.JobWithOptions("crawling_job", work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).PerformCrawling)

	pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	pool.Stop()
}
