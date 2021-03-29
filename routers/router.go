package routers

import (
	"google-scraper/controllers"

	"github.com/beego/beego/v2/server/web"
	apiv1 "google-scraper/controllers/api/v1"
)

func init() {
	web.Router("/", &controllers.Dashboard{}, "get:New")
	web.Router("/search", &controllers.Search{}, "post:Create")
	web.Router("/register", &controllers.Registration{}, "get:New;post:Create")
	web.Router("/login", &controllers.Session{}, "get:New;post:Create")
	web.Router("/logout", &controllers.Session{}, "get:Delete")
	web.Router("/keyword/:id:int", &controllers.KeywordController{}, "get:Show")
	web.Router("/keyword/:id:int/render-html", &controllers.KeywordController{}, "get:RenderHtml")

	// API
	// init namespace
	ns := web.NewNamespace("/api/v1",
		web.NSRouter("/me", &apiv1.Me{}),
	)

	// register namespace
	web.AddNamespace(ns)
}
