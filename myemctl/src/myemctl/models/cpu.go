package models

import (
	"github.com/astaxie/beego/orm"
)
type Cpu struct{
	Id 		int
	Addtime		int
	Hostname	string
	Cpu_user_pct	int
	Cpu_nice_pct	int
	Cpu_sys_pct	int
	Cpu_idle_pct	int
	Cpu_iowait_pct	int
	Cpu_irq_pct	int
	Cpu_softirq_pct int
	Cpu_all_irq_pct int
	Cpu_total_pct	int
}

func (u *Cpu) TableName() string{
	return "sys_cpuinfo"
}

func init(){
	orm.RegisterModel(new(Cpu))
}
