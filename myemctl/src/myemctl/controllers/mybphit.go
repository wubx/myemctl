package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myemctl/models"
	"fmt"
	"os"
)

type MyBpHitController struct {
	beego.Controller
}
/*
type  genernal interface{}
type  Kv  [2]genernal
*/
type  MyHit struct{
	Hit	[]Kv
}

func (this *MyBpHitController) Get() {

	o := orm.NewOrm()
	mybphit := new(models.MyBpHit)

	var hits  []*models.MyBpHit
	hostname, _ := os.Hostname()
	num, err := o.QueryTable(mybphit).Filter("hostname", hostname).All(&hits)
	if err != nil{
		fmt.Println(err.Error())
	}
	var  bphit []Kv = make([]Kv, num)

	//fmt.Println(num)
	for i := 0 ; i < len(hits) ; i++{
		bphit[i][0] = hits[i].Addtime*1000
		bphit[i][1] = hits[i].I_bp_hit
	}
	var    bphit_a MyHit
	bphit_a.Hit = bphit
	this.Data["json"] = &bphit_a
	this.ServeJson()

}

