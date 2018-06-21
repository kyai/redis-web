package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) JSON(code int, msg string, data interface{}) {
	var jdata = make(map[string]interface{})
	jdata["Code"] = code
	jdata["Msg"] = msg
	jdata["Data"] = data
	c.Data["json"] = jdata
	c.ServeJSON()
	return
}
