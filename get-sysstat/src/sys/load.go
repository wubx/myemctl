package sys
import (
	"fmt"
	"os"
	"comm"
	"strings"
	"bufio"
	"bytes"
	"encoding/csv"
	"strconv"
)

func get_load() (load []string){
	file, err := os.Open("/proc/loadavg")
	defer file.Close()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	r := bufio.NewReader(file)
	line , err := r.ReadString('\n')
	load =strings.Fields(line)
	return 
}

func OutputLoad(t int,mdb *comm.DBconn ){
	//interval := 1
	loadFile := "./output/load.csv"
	load_fd , err :=  os.OpenFile(loadFile,os.O_RDWR|os.O_APPEND, 0660)
	defer load_fd.Close()
	if err != nil{
		load_fd , err =  os.Create(loadFile)
	}

	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	tmp := get_load()
	tmp1 := make([]string,6)
	tmp1[0] = strconv.Itoa(t)
	for  i := 0; i<len(tmp) ;i++{
		tmp1[i+1] = tmp[i]
	}
	r2.Write(tmp1)
	r2.Flush()
	load_fd.WriteString(buf.String())
	
	hostname, _ := os.Hostname()
	dbx := mdb.GetMySQLconn()
	defer mdb.PutMySQL(dbx)
	stmt, err := dbx.Prepare("insert into sys_load(addtime,hostname,la1, la5, la15,pid_and_num, run_pid)values(?,?,?,?,?,?,?)")
	if  err != nil{
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	stmt.Exec(t, hostname,tmp[0], tmp[1], tmp[2], tmp[3], tmp[4])
}
