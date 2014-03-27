package main

import (
	_ "myemctl/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init(){
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpassword")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqldb   := beego.AppConfig.String("mysqldb")
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	//fmt.Println(mysqluser, mysqlpass, mysqlurls, mysqlport, mysqldb)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",mysqluser, mysqlpass, mysqlurls, mysqlport, mysqldb))
}

func main() {
	beego.Run()
}

