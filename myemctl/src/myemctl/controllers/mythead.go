package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myemctl/models"
	"fmt"
	"os"
)

type MyThdController struct {
	beego.Controller
}
/*
type  genernal interface{}
type  Kv  [2]genernal
*/
type  MyThd struct{
	ThdRun	[]Kv
	ThdConn []Kv
	ThdCache []Kv
	ThdMaxConn []Kv
	ThdCreated	[]Kv
	ThdAborted	[]Kv
}

func (this *MyThdController) Get() {

	o := orm.NewOrm()
	thd := new(models.MyThd)

	var thds []*models.MyThd
	hostname, _ := os.Hostname()
	num, err := o.QueryTable(thd).Filter("hostname", hostname).All(&thds)
	if err != nil{
		fmt.Println(err.Error())
	}
	var  ThdRun []Kv = make([]Kv, num)
	var  ThdConn []Kv = make([]Kv, num)
	var  ThdCache []Kv = make([]Kv, num)
	var  ThdMaxConn []Kv = make([]Kv, num)
	var  ThdCreated []Kv = make([]Kv, num)
	var  ThdAborted []Kv = make([]Kv, num)

	//fmt.Println(num)
	for i := 0 ; i < len(thds) ; i++{
		ThdRun[i][0] = thds[i].Addtime*1000
		ThdRun[i][1] = thds[i].T_run

		ThdConn[i][0] = thds[i].Addtime*1000
		ThdConn[i][1] = thds[i].T_conned
		
		ThdCache[i][0] = thds[i].Addtime*1000
		ThdCache[i][1] =  thds[i].T_cached

		ThdMaxConn[i][0] = thds[i].Addtime*1000
		ThdMaxConn[i][1] = thds[i].T_max_conn

		ThdCreated[i][0] = thds[i].Addtime*1000
		ThdCreated[i][1] = thds[i].T_thread_created

		ThdAborted[i][0] = thds[i].Addtime*1000
		ThdAborted[i][1] = thds[i].T_aborted_conn
	}
	var   mythd_a  MyThd
	mythd_a.ThdRun = ThdRun
	mythd_a.ThdConn = ThdConn
	mythd_a.ThdCache = ThdCache
	mythd_a.ThdMaxConn = ThdMaxConn
	mythd_a.ThdCreated = ThdCreated
	mythd_a.ThdAborted = ThdAborted
	this.Data["json"] = &mythd_a

	this.ServeJson()

}

