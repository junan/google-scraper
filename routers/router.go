package routers

import (
	"google-scraper/controllers"

	 "github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.Dashboard{})
	web.Router("/register", &controllers.Registration{})
	web.Router("/login", &controllers.Session{})
	web.Router("/logout", &controllers.Session{}, "get:Delete")
}
