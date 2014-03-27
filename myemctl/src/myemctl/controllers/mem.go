package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myemctl/models"
	"fmt"
	"os"
)

type MemController struct {
	beego.Controller
}
/*
type  genernal interface{}
type  Kv  [2]genernal
*/
type  MemPer struct{
	Use	 []Kv
	Free 	[]Kv
	Total	[]Kv
}

func (this *MemController) Get() {

	o := orm.NewOrm()
	mem := new(models.Mem)

	var mems []*models.Mem
	hostname, _ := os.Hostname()
	num, err := o.QueryTable(mem).Filter("hostname", hostname).All(&mems)
	if err != nil{
		fmt.Println(err.Error())
	}
	var  Use []Kv = make([]Kv, num)
	var  Free []Kv = make([]Kv, num)
	var  Total []Kv = make([]Kv, num)
	//fmt.Println(num)
	for i := 0 ; i < len(mems) ; i++{
		Use[i][0] = mems[i].Addtime*1000
		Use[i][1] =  mems[i].Mem_total - mems[i].Mem_free

		Free[i][0] = mems[i].Addtime*1000
		Free[i][1] =  mems[i].Mem_free
		
		Total[i][0] = mems[i].Addtime*1000
		Total[i][1] = mems[i].Mem_total
	}
	var   mem_s  MemPer
	mem_s.Use  = Use
	mem_s.Free = Free
	mem_s.Total = Total
	this.Data["json"] = &mem_s

	this.ServeJson()

}

