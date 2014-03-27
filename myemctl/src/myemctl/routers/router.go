package routers

import (
	"myemctl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/cpu", &controllers.CpuController{})
    beego.Router("/load", &controllers.LoadController{})
    beego.Router("/mem", &controllers.MemController{})
    beego.Router("/thd", &controllers.MyThdController{})
    beego.Router("/active", &controllers.MyActiveController{})
    beego.Router("/bphit", &controllers.MyBpHitController{})
}
