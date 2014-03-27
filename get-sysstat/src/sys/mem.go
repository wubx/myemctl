package sys
import (
	"os/exec"
	"fmt"
	"strings"
	"bytes"
	"os"
	"comm"
)

func OutputMem(t int, mdb *comm.DBconn){
	memfile := "./output/mem.csv"
	mem_fd, err := os.OpenFile(memfile,os.O_RDWR|os.O_APPEND, 0660)
	defer mem_fd.Close()
	if err != nil{
		mem_fd, err = os.Create(memfile)
	}
	cmd := exec.Command("bash", "-c", "cat /proc/meminfo |grep MemTotal -A 1|awk '{print $2}'|xargs|awk '{printf \"%d %d\",$1/1024,$2/1024}'")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	tmp :=  strings.Fields(out.String())
	fmt.Fprintf(mem_fd,"%d,%s,%s\n",t,tmp[0],tmp[1])

	hostname , _ := os.Hostname()
	dbx := mdb.GetMySQLconn()
	defer mdb.PutMySQL(dbx)
	stmt, err := dbx.Prepare("insert into sys_mem(addtime,hostname,mem_total, mem_free) values(?, ?, ?, ?)")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	stmt.Exec(t, hostname, tmp[0], tmp[1])
}

