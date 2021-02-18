package routers

import (
	"google-scraper/controllers"

	 "github.com/beego/beego/v2/server/web"
)

func init() {
	// TODO: replace it with real controller
	web.Router("/", &controllers.Home{})
	web.Router("/register", &controllers.Registration{})
	web.Router("/login", &controllers.Session{})
}
