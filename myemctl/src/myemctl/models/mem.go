package models

import (
	"github.com/astaxie/beego/orm"
)
type Mem struct{
	Id		int
	Addtime		int
	Hostname	string
	Mem_total	int
	Mem_free	int

}

func (u *Mem) TableName() string{
	return "sys_mem"
}

func init(){
	orm.RegisterModel(new(Mem))
}
