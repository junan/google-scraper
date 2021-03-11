package main

import (
	"os"
	"os/signal"
	"syscall"

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

	pool.Middleware((*jobs.Context).PerformLater)

	pool.JobWithOptions(conf.GetString("scraperJobName"), work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).PerformScrape)

	pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	pool.Stop()
}
