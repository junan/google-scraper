package initializers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"

	"google-scraper/models"
)

const (
	layoutUS = "January 2, 2006"
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

func DisplayStatus(k models.Keyword) string {
	if k.SearchCompleted {
		return "Yes"
	} else {
		return "No"
	}
}

func DisplayFormattedCreatedDate(k models.Keyword) string {
	return k.CreatedAt.Format(layoutUS)
}
