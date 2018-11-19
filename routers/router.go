package routers

import (
	"github.com/astaxie/beego"
	"github.com/palagend/ceego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/cmd", &controllers.CommandController{}, "get:Echo")
}
