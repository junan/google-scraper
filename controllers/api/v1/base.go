package apiv1

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/jsonapi"
)

type baseAPIController struct {
	web.Controller
}

func (c *baseAPIController) serveJSON(data interface{}) {
	response, err := jsonapi.Marshal(data)
	if err != nil {
		c.renderError(err)
		return
	}

	c.Data["json"] = response
	err = c.ServeJSON()
	if err != nil {
		c.renderError(err)
	}
}

func (c *baseAPIController) renderError(err error) {
	err = jsonapi.MarshalErrors(c.Ctx.ResponseWriter, []*jsonapi.ErrorObject{{
		Detail: err.Error(),
	}})
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}
