package models

import (
	"github.com/astaxie/beego/orm"
)

type MyThd struct {
	Id               int
	Addtime          int
	Hostname         string
	Port             int
	T_run            int
	T_conned         int
	T_cached         int
	T_max_conn       int
	T_thread_created int
	T_aborted_conn   int
}

func (u *MyThd) TableName() string {
	return "my_thread"
}

func init() {
	orm.RegisterModel(new(MyThd))
}
