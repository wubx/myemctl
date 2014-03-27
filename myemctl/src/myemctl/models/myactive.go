package models

import (
	"github.com/astaxie/beego/orm"
)
type MyActive struct{
	Id		int
	Addtime		int
	Hostname	string
	Sselect		int
	Uupdate		int
	Ddelete		int
	Iinsert		int
	Ccall		int
}

func (u *MyActive) TableName() string{
	return "my_active"
}

func init(){
	orm.RegisterModel(new(MyActive))
}
