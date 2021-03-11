package main

import (
	"os"
	"os/signal"
	"syscall"

	"google-scraper/database"
	_ "google-scraper/initializers"
	"google-scraper/worker/jobs"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/beego/beego/v2/core/logs"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial:      database.GetRedisUrl,
}

func init() {
}

func main() {
	logs.Error("Main worker is running: ")
	pool := work.NewWorkerPool(jobs.Context{}, 5, "google_scraper_queue", redisPool)

	pool.JobWithOptions("crawling_job", work.JobOptions{MaxFails: jobs.MaxFails}, (*jobs.Context).PerformCrawling)

	pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	pool.Stop()
}
