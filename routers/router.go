package routers

import (
	"google-scraper/controllers"
	apiv1 "google-scraper/controllers/api/v1"

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
	web.Router("/oauth-client", &controllers.OauthClient{}, "get:New;post:Create")

	// API
	// init namespace
	ns := web.NewNamespace("/api/v1",
		web.NSRouter("/health-check", &apiv1.HealthCheck{}),
		web.NSRouter("/token", &apiv1.Token{}, "post:Create"),
		web.NSRouter("/revoke", &apiv1.Token{}, "post:Revoke"),
		web.NSRouter("/keyword", &apiv1.Keyword{}, "post:Upload"),
	)

	// register namespace
	web.AddNamespace(ns)
}
