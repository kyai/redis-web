package routers

import (
	"redis-web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/setup", &controllers.SetupController{}, "get:Index")
	beego.Router("/api/CheckConn", &controllers.SetupController{}, "post:CheckConn")
}
