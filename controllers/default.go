package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.meeeeeeee！！！！！！"
	c.Data["Email"] = "bmjbmj@gmail.com"
	c.TplName = "index.tpl"
}
