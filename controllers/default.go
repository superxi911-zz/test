package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "try opo"
	c.Data["Email"] = "agaihghiinjjjj@gmail.com"
	c.TplName = "index.tpl"
}
