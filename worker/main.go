package main

import (
	"os"
	"os/signal"
	"syscall"

	"google-scraper/initializers"
	"google-scraper/database"
	"google-scraper/worker/jobs"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial:      database.GetRedisUrl,
}

func init() {
	bootstrap.SetUp()
}

func main() {
	pool := work.NewWorkerPool(jobs.Context{}, 5, "google_scraper_queue", redisPool)

	//pool.Middleware((*jobs.Context).PerformCrawling)

	pool.JobWithOptions("crawling_job", work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).PerformCrawling)

	pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	pool.Stop()
}
