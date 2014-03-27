package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myemctl/models"
	"os"
)

type MyActiveController struct {
	beego.Controller
}

/*
type  genernal interface{}
type  Kv  [2]genernal
*/
type MyAct struct {
	Sselect []Kv
	Uupdate []Kv
	Ddelete []Kv
	Iinsert []Kv
	Ccall   []Kv
}

func (this *MyActiveController) Get() {

	o := orm.NewOrm()
	myactive := new(models.MyActive)

	var acts []*models.MyActive
	hostname, _ := os.Hostname()
	num, err := o.QueryTable(myactive).Filter("hostname", hostname).All(&acts)
	if err != nil {
		fmt.Println(err.Error())
	}
	var Sselect []Kv = make([]Kv, num)
	var Uupdate []Kv = make([]Kv, num)
	var Ddelete []Kv = make([]Kv, num)
	var Iinsert []Kv = make([]Kv, num)
	var Ccall []Kv = make([]Kv, num)

	//fmt.Println(num)
	for i := 0; i < len(acts); i++ {
		Sselect[i][0] = acts[i].Addtime * 1000
		Sselect[i][1] = acts[i].Sselect

		Uupdate[i][0] = acts[i].Addtime * 1000
		Uupdate[i][1] = acts[i].Uupdate

		Ddelete[i][0] = acts[i].Addtime * 1000
		Ddelete[i][1] = acts[i].Ddelete

		Iinsert[i][0] = acts[i].Addtime * 1000
		Iinsert[i][1] = acts[i].Iinsert

		Ccall[i][0] = acts[i].Addtime * 1000
		Ccall[i][1] = acts[i].Ccall
	}
	var myact_a MyAct
	myact_a.Sselect = Sselect
	myact_a.Uupdate = Uupdate
	myact_a.Ddelete = Ddelete
	myact_a.Iinsert = Iinsert
	myact_a.Ccall = Ccall
	this.Data["json"] = &myact_a
	this.ServeJson()

}
