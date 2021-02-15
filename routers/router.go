package routers

import (
	"google-scraper/controllers"

	 "github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/register", &controllers.Registration{})
}
