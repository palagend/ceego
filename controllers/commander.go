package controllers

import (
	"github.com/astaxie/beego"
)

// CommandController controller to handle command
type CommandController struct {
	beego.Controller
}

// Echo echo message
func (cc *CommandController) Echo() {
	name := cc.Input().Get("name")
	cc.Ctx.WriteString("How are you, " + name)
}
