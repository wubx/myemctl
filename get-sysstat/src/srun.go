package main

import (
	"sys"
	//	"runtime"
	"comm"
	"fmt"
	"mysql"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//读取配置
	c, err := comm.ReadConfigFile("stat.conf")
	if err != nil {
		fmt.Println(err)
	}
	node, err := c.GetString("default", "node")
	if err != nil {
		fmt.Println(err)
	}
	mount_node, err := c.GetString("default", "mount_node")
	if err != nil {
		fmt.Println(err)
	}

	val_1, _ := c.GetString("default", "interval")
	inter_i, _ := strconv.Atoi(val_1)

	user, _ := c.GetString("mysql", "user")
	password, _ := c.GetString("mysql", "password")
	port, _ := c.GetString("mysql", "port")

	//mondb
	mon_host, _ := c.GetString("mondb", "host")
	mon_port, _ := c.GetString("mondb", "port")
	mon_user, _ := c.GetString("mondb", "user")
	mon_password, _ := c.GetString("mondb", "password")
	mdb := &comm.DBconn{mon_user, mon_password, mon_host, mon_port, "mondb"}
	mysql.InitStr(mdb)
	fmt.Println("开始收集信息...")
	outputdir := "output"
	os.RemoveAll(outputdir)
	os.Mkdir(outputdir, 666)
	//time.Sleep(2 * time.Second)
	for {
		t := int(time.Now().Unix())
		go sys.OutputCpu(t, mdb)
		sys.OutputLoad(t, mdb)
		ndev := strings.Split(node, ",")
		go sys.OutputDiskioByArray(t, ndev, mdb)
		//disk space info
		mdev := strings.Split(mount_node, ",")
		for j := range mdev {
			sys.OutputDisksp(t, mdev[j], mdb)
		}
		go sys.OutputMem(t, mdb)
		go mysql.OutputMySQL(user, password, port, t, mdb)
		//	runtime.Gosched()
		time.Sleep(time.Duration(inter_i) * time.Second)
	}
}
