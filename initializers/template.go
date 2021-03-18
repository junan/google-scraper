package initializers

import (
	. "google-scraper/helpers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)


func init() {
	err := web.AddFuncMap("displayStatus", DisplayStatus)
	if err != nil {
		logs.Error("Registering displayStatus function to the template failed: ", err)
	}

	err = web.AddFuncMap("displayFormattedCreatedDate", DisplayFormattedCreatedDate)
	if err != nil {
		logs.Error("Registering displayFormattedCreatedDate function to the template failed: ", err)
	}
}

