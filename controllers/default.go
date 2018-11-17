package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "ceego.top"
	c.Data["Email"] = "palagend@aliyun.com"
	c.TplName = "index.tpl"
}
