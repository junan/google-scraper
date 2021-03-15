package helpers

import (
	"google-scraper/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/iancoleman/strcase"
)

func SetDataAttributes(c *web.Controller, user *models.User) {
	controllerName, actionName := c.GetControllerAndAction()

	c.Layout = "layouts/default.html"
	c.Data["ControllerName"] = strcase.ToSnake(controllerName)
	c.Data["ActionName"] = strcase.ToSnake(actionName)
	c.Data["CurrentUser"] = user
	c.Data["AlertMap"] = map[string]string{
		"success": "success",
		"error":   "danger",
	}
}
