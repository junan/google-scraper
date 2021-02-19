package routers

import (
	"google-scraper/controllers"

	 "github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.Dashboard{}, "get:New")
	web.Router("/register", &controllers.Registration{}, "get:New")
	web.Router("/register", &controllers.Registration{}, "get:Create")
	web.Router("/login", &controllers.Session{}, "get:New")
	web.Router("/login", &controllers.Session{}, "post:Create")
	web.Router("/logout", &controllers.Session{}, "get:Delete")
}
