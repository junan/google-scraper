package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "google-scraper/initializers"
	"google-scraper/database"
	"google-scraper/worker/jobs"


	"github.com/gocraft/work"
)

func init() {
}

func main() {
	pool := work.NewWorkerPool(jobs.Context{}, 5, "google_scraper", database.GetRedisPool())
	pool.JobWithOptions("crawling_job", work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).PerformCrawling)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	// Stop the pool
	pool.Stop()
}
