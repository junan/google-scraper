package initializers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func LoadAppConfig() {
	err := web.LoadAppConfig("ini", "../conf/app.conf")

	if err != nil {
		logs.Error("Error is while loading the config: ", err)
	}
}
