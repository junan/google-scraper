package main

import (
	_ "google-scraper/initializers"
	_ "google-scraper/routers"

	_ "github.com/beego/beego/v2/server/web/session/redis"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
