package helpers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/iancoleman/strcase"
)

func SetDataAttributes(c *web.Controller) {
	controllerName, actionName := c.GetControllerAndAction()

	c.Data["ControllerName"] = strcase.ToSnake(controllerName)
	c.Data["ActionName"] =  strcase.ToSnake(actionName)
}
