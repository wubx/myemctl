package models

import (
	"github.com/astaxie/beego/orm"
)
type Load struct{
	Id		int
	Addtime		int
	Hostname	string
	La1	float32
	La5	float32
	La15	float32
	Pid_and_num  string
	Run_pid string
}

func (u *Load) TableName() string{
	return "sys_load"
}

func init(){
	orm.RegisterModel(new(Load))
}
