package main

import (
	. "google-scraper/helpers"
	_ "google-scraper/initializers"
	_ "google-scraper/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.AddFuncMap("displayStatus", DisplayStatus)
	web.AddFuncMap("displayFormattedCreatedDate", DisplayFormattedCreatedDate)

	web.Run()
}
