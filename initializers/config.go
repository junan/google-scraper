package initializers

import (
	"fmt"

	helpers "google-scraper/tests"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func LoadAppConfig() {
	configPath := fmt.Sprintf("%s/conf/app.conf", helpers.AppRootDir(1))
	err := web.LoadAppConfig("ini", configPath)

	if err != nil {
		logs.Error("Error is while loading the config: ", err)
	}
}
