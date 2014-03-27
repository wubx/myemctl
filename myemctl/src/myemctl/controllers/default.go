package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Data["Website"] = "www.mysqlsupport.cn"
	this.Data["Email"] = "wubingxi@gmail.com"
	this.TplNames = "index.tpl"
}
