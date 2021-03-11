package controllers

import (
	"log"
	"os"
	"os/signal"

	"github.com/beego/beego/v2/server/web"
	"github.com/gocraft/work"

	. "google-scraper/initializers"
	"google-scraper/models"
)

type Context struct {
	customerID int64
}

type Dashboard struct {
	baseController
}

func (c *Dashboard) New() {
	web.ReadFromRequest(&c.Controller)

	c.setAttributes()

	// Make an enqueuer with a particular namespace
	var enqueuer = work.NewEnqueuer("google_scraper", RedisPool)

	_, err := enqueuer.Enqueue("perform_search_job", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})
	if err != nil {
		log.Fatal(err)
	}

	// Make a new pool. Arguments:
	// Context{} is a struct that will be the context for the request.
	// 10 is the max concurrency
	// "my_app_namespace" is the Redis namespace
	// redisPool is a Redis pool
	pool := work.NewWorkerPool(Context{}, 10, "my_app_namespace", RedisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).FindKeyword)

	// Map the name of jobs to handler functions
	pool.Job("send_email", (*Context).SendEmail)

	// Customize options:
	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func (c *Dashboard) setAttributes() {
	c.TplName = "dashboard/new.html"
}

func FindKeyword(ID int64) models.Keyword {

}

