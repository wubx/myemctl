package models

import (
	"github.com/astaxie/beego/orm"
)
type MyBpHit struct{
	Id		int
	Addtime		int
	Hostname	string
	I_bp_hit	int

}

func (u *MyBpHit) TableName() string{
	return "my_innodb_status"
}

func init(){
	orm.RegisterModel(new(MyBpHit))
}
