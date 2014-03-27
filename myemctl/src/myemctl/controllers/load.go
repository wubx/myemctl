package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myemctl/models"
	"os"
)

type LoadController struct {
	beego.Controller
}

/*
type  genernal interface{}
type  Kv  [2]genernal
*/
type LoadPer struct {
	La1  []Kv
	La5  []Kv
	La15 []Kv
}

func (this *LoadController) Get() {

	o := orm.NewOrm()
	load := new(models.Load)

	var loads []*models.Load
	hostname, _ := os.Hostname()
	num, err := o.QueryTable(load).Filter("hostname", hostname).All(&loads)
	if err != nil {
		fmt.Println(err.Error())
	}
	var La1 []Kv = make([]Kv, num)
	var La5 []Kv = make([]Kv, num)
	var La15 []Kv = make([]Kv, num)
	//fmt.Println(num)
	for i := 0; i < len(loads); i++ {
		La1[i][0] = loads[i].Addtime * 1000
		La1[i][1] = loads[i].La1
		La5[i][0] = loads[i].Addtime * 1000
		La5[i][1] = loads[i].La5
		La15[i][0] = loads[i].Addtime * 1000
		La15[i][1] = loads[i].La15
	}
	var load_p LoadPer
	load_p.La1 = La1
	load_p.La5 = La5
	load_p.La15 = La15
	this.Data["json"] = &load_p

	this.ServeJson()

}
