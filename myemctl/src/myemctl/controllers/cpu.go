package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myemctl/models"
	"os"
)

type CpuController struct {
	beego.Controller
}

/*
type Kv struct{
	Addtime int
	Per	float32
}
*/
type genernal interface{}
type Kv [2]genernal

type CpuPer struct {
	CpuTotal  []Kv
	CpuIdle   []Kv
	CpuSys    []Kv
	CpuUser   []Kv
	CpuIowait []Kv
	CpuIrq    []Kv
}

func (this *CpuController) Get() {
	/*
		mysqluser := beego.AppConfig.String("mysqluser")
		mysqlpass := beego.AppConfig.String("mysqlpass")
		mysqlurls := beego.AppConfig.String("mysqlurls")
		mysqlport := beego.AppConfig.String("mysqlport")
		mysqldb   := beego.AppConfig.String("mysqldb")
		mdb := comm.DBconn{mysqluser, mysqlpass, mysqlurls, mysqlport, mysqldb}
		this.Data["json"] = mdb
		this.ServeJson()
	*/
	o := orm.NewOrm()
	cpu := new(models.Cpu)

	var cpus []*models.Cpu
	hostname, _ := os.Hostname()
	num, err := o.QueryTable(cpu).Filter("hostname", hostname).All(&cpus)
	if err != nil {
		fmt.Println(err.Error())
	}

	var CpuTotal []Kv = make([]Kv, num)
	var CpuIdle []Kv = make([]Kv, num)
	var CpuSys []Kv = make([]Kv, num)
	var CpuUser []Kv = make([]Kv, num)
	var CpuIowait []Kv = make([]Kv, num)
	var CpuIrq []Kv = make([]Kv, num)

	for i := 0; i < len(cpus); i++ {
		//fmt.Println("Cpu_total_pct: ",cpus[i].Addtime, cpus[i].Cpu_total_pct)
		CpuTotal[i][0] = cpus[i].Addtime * 1000
		CpuTotal[i][1] = float32(cpus[i].Cpu_total_pct) / 100

		//fmt.Println("Cpu_idle_pct ", cpus[i].Addtime, cpus[i].Cpu_idle_pct)
		CpuIdle[i][0] = cpus[i].Addtime * 1000
		CpuIdle[i][1] = float32(cpus[i].Cpu_idle_pct) / 100

		//fmt.Println("Cpu_sys_pct ", cpus[i].Addtime, cpus[i].Cpu_sys_pct)
		CpuSys[i][0] = cpus[i].Addtime * 1000
		CpuSys[i][1] = float32(cpus[i].Cpu_sys_pct) / 100

		//fmt.Println("Cpu_user_pct ", cpus[i].Addtime, cpus[i].Cpu_user_pct)
		CpuUser[i][0] = cpus[i].Addtime * 1000
		CpuUser[i][1] = float32(cpus[i].Cpu_user_pct) / 100

		//fmt.Println("Cpu_iowait_pct ", cpus[i].Addtime, cpus[i].Cpu_iowait_pct)
		CpuIowait[i][0] = cpus[i].Addtime * 1000
		CpuIowait[i][1] = float32(cpus[i].Cpu_iowait_pct) / 100

		//fmt.Println("Cpu_irq_pct ", cpus[i].Addtime, cpus[i].Cpu_all_irq_pct)
		CpuIrq[i][0] = cpus[i].Addtime * 1000
		CpuIrq[i][1] = float32(cpus[i].Cpu_all_irq_pct) / 100
	}
	var cpu_pct CpuPer
	cpu_pct.CpuTotal = CpuTotal
	cpu_pct.CpuIdle = CpuIdle
	cpu_pct.CpuSys = CpuSys
	cpu_pct.CpuUser = CpuUser
	cpu_pct.CpuIowait = CpuIowait
	cpu_pct.CpuIrq = CpuIrq
	this.Data["json"] = &cpu_pct

	this.ServeJson()

}
