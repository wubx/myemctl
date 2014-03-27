package sys
import (
	"bytes"
	"io"
	"time"
	"os"
	"os/exec"
	"strings"
	"comm"
	"fmt"
)

func get_diskioinfo(dev string)( tmp []string) {
	cmd1 := exec.Command("cat", "/proc/diskstats")	
	cmd2 := exec.Command("grep", "-w",dev)	
	r, w := io.Pipe()
	cmd1.Stdout= w
	cmd2.Stdin = r

	var out bytes.Buffer
	cmd2.Stdout = &out
	cmd1.Start()
	cmd2.Start()
	cmd1.Wait()
	w.Close()
	cmd2.Wait()
//	fmt.Printf("%q\n",strings.Fields(out.String()))
	tmp = strings.Fields(out.String())[2:]
//	fmt.Printf("%q\n", tmp1)
	return
}

func OutputDiskioByArray(t int, ndev []string, mdb *comm.DBconn){
	for i :=  range(ndev){
		OutputDiskio(t, ndev[i], mdb)
		//fmt.Println(ndev[i])
	}
}
func OutputDiskio(t int, dev string, mdb *comm.DBconn){
	tmp1 :=  get_diskioinfo(dev)
	diskiofile :="./output/diskioinfo.cvs"
	io_fd, err :=  os.OpenFile(diskiofile,os.O_RDWR|os.O_APPEND, 0660)
	if err != nil{
		io_fd, err =  os.Create(diskiofile)
	}
	defer io_fd.Close()

	time.Sleep(5*time.Second)
	tmp2 := get_diskioinfo(dev)
	//addtime, dev ,reads,rd_mrg,rd_sectors,ms_reading,writes,wr_mrg,wr_sectors,ms_writing,ms_doing_io
	//每秒读多少次
	dev_read_pers :=  int(comm.Str_trade(tmp2[0],tmp1[0])/5)
	//每秒有多次读合并
	dev_readmerge_pers := int(comm.Str_trade(tmp2[1],tmp1[1])/5)
	//每秒读多少个扇区
	dev_readsector_pers := int(comm.Str_trade(tmp2[2],tmp1[2])/5)
	//每秒读用了多少时间
	dev_read_use_pers := int(comm.Str_trade(tmp2[3],tmp1[3])/5)
	//每秒有多个写
	dev_write_pers  := int(comm.Str_trade(tmp2[4],tmp1[4])/5)
	//每秒有多少个写合并
	dev_writemerge_pers := int(comm.Str_trade(tmp2[5],tmp1[5])/5)
	//每秒写了多少个扇区
	dev_writesector_pers := int(comm.Str_trade(tmp2[6],tmp1[6])/5)
	//每秒用在写上的时间
	dev_write_use_pers := int(comm.Str_trade(tmp2[7],tmp1[7])/5)
	//每秒有多少个io
	dev_io_pers :=  int(comm.Str_trade(tmp2[9],tmp1[9])/5)
	fmt.Fprintf(io_fd, "%d,%s,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", t, dev, dev_read_pers, dev_readmerge_pers, dev_readsector_pers, dev_read_use_pers, dev_write_pers, dev_writemerge_pers, dev_writesector_pers, dev_write_use_pers, dev_io_pers)
	tmp1 = tmp2
	
	hostname , _ := os.Hostname()

	dbx := mdb.GetMySQLconn()
	defer mdb.PutMySQL(dbx)
	//stmt , err := dbx.Prepare("insert into sys_diskioinfo(addtime, hostname, dev_name, dev_read_pers, dev_readmerge_pers) values(?, ?, ?, ?, ?)")
	//x, dev_readsector_pers, dev_read_use_pers, dev_write_pers, dev_writemerge_pers, dev_writesector_pers,  dev_write_use_pers,dev_io_pers) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	stmt , err := dbx.Prepare("insert into sys_diskioinfo(addtime, hostname, dev_name, dev_read_pers, dev_readmerge_pers, dev_readsector_pers, dev_read_use_pers, dev_write_pers, dev_writemerge_pers, dev_writesector_pers,  dev_write_use_pers,dev_io_pers) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	//_, err = stmt.Exec(t, hostname, dev, dev_read_pers, dev_readmerge_pers)
	_, err = stmt.Exec(t, hostname, dev, dev_read_pers, dev_readmerge_pers, dev_readsector_pers, dev_read_use_pers, dev_write_pers, dev_writemerge_pers, dev_writesector_pers, dev_write_use_pers, dev_io_pers)
	//_, err = stmt.Exec(t, hostname, dev, s[2], s[3], s[4],s[5],s[6],s[7],s[8],s[9],s[10])
	if err != nil{
		fmt.Println(err.Error())
		return 
	}
}
