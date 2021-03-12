package tests

import (
	"google-scraper/database"

	"github.com/beego/beego/v2/core/logs"
)

func DeleteRedisJobs(namespace, jobName string) {
	connection := database.GetRedisPool().Get()
	defer connection.Close()

	_, err := connection.Do("DEL", redisKeyJobs(namespace, jobName))

	if err != nil {
		logs.Error("Deleting redid jobs failed: ", err)
	}
}

func redisKeyJobs(namespace, jobName string) string {
	return redisKeyJobsPrefix(namespace) + jobName
}

func redisKeyJobsPrefix(namespace string) string {
	return redisNamespacePrefix(namespace) + "jobs:"
}

func redisNamespacePrefix(namespace string) string {
	l := len(namespace)
	if (l > 0) && (namespace[l-1] != ':') {
		namespace = namespace + ":"
	}
	return namespace
}
