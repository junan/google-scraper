package initializers

import (
	"github.com/beego/beego/v2/server/web"

	"google-scraper/models"
)

const (
	layoutUS = "January 2, 2006"
)

func init() {
	web.AddFuncMap("displayStatus", DisplayStatus)
	web.AddFuncMap("displayFormattedCreatedDate", DisplayFormattedCreatedDate)
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
