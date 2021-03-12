package tests

import (
	"google-scraper/database"

	"github.com/gocraft/work"
)

func GetWorkerClient(workerNamespace string) *work.Client {
	return work.NewClient(workerNamespace, database.GetRedisPool())
}
