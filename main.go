package main

import (
	_ "google-scraper/initializers"
	_ "google-scraper/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
