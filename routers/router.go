package routers

import (
	"google-scraper/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.Dashboard{}, "get:New")
	web.Router("/search", &controllers.Search{}, "post:Create")
	web.Router("/register", &controllers.Registration{}, "get:New;post:Create")
	web.Router("/login", &controllers.Session{}, "get:New;post:Create")
	web.Router("/logout", &controllers.Session{}, "get:Delete")
	web.Router("/keyword/:id:int", &controllers.KeywordController{}, "get:Show")
	web.Router("/keyword/:id:int/render-html", &controllers.KeywordController{}, "get:RenderHtml")
}
