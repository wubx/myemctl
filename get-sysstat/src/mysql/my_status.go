package mysql
import (
	"fmt"
	"database/sql"
	"strconv"
	"time"
	"os"
	"comm"
//	"encoding/binary"
	_ "github.com/go-sql-driver/mysql"
)
var  tmp1 map[string]int

func  get_mystr(user,password,port string) (tmp_str map[string]int) {
	tmp_str = make(map[string] int)
	db ,err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/mysql", user,password,port))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, er := db.Query("SHOW GLOBAL STATUS")
	if  er != nil{
		fmt.Println(er)
		return
	}
	cols, _ := rows.Columns()

	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{} , len(values))

	for i := range values{
		scanArgs[i] = &values[i]
	}

	for rows.Next(){
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		key , err := strconv.Atoi(fmt.Sprintf("%s",values[1]))
		//key , err := binary.ReadVarint(values[1])
		if  err != nil{
			key = 0
		}
		tmp_str[fmt.Sprintf("%s", values[0])] = key
	}
	return 
}

func my_active(tmp1, tmp2 map[string]int, addtime int, port string, mdb *comm.DBconn) {
	
	hostname , _ := os.Hostname()
	dbx := mdb.GetMySQLconn()
	defer mdb.PutMySQL(dbx)

	//数据库活跃计算 my_active
	t := tmp2["Uptime"] - tmp1["Uptime"]
	sselect := (tmp2["Com_select"] - tmp1["Com_select"])/t
	uupdate := (tmp2["Com_update"] + tmp2["Com_update_multi"] - tmp1["Com_update"] - tmp1["Com_update_multi"])/t
	ddelete := (tmp2["Com_delete"] + tmp2["Com_delete_multi"] - tmp1["Com_delete"] - tmp1["Com_delete_multi"])/t
	iinsert := (tmp2["Com_insert"] + tmp2["Com_insert_select"] - tmp1["Com_insert"] - tmp1["Com_insert_select"])/t
	ccall   := (tmp2["Com_call_procedure"] - tmp1["Com_call_procedure"])/t

	q_row_insert := (tmp2["Handler_write"] - tmp1["Handler_write"])/t
	q_row_update  := (tmp2["Handler_update"] - tmp1["Handler_update"])/t
	q_row_delete  := (tmp2["Handler_delete"] - tmp1["Handler_delete"])/t
	my_fd , err := os.OpenFile("./output/my_active.csv",os.O_RDWR|os.O_APPEND, 0660)
	defer my_fd.Close()
	if err != nil{
		my_fd, _ = os.Create("./output/my_active.csv")
	}
	fmt.Fprintf(my_fd, "%d,%d,%d,%d,%d,%d,%d,%d,%d\n",addtime, sselect,uupdate, ddelete, iinsert, ccall,q_row_insert,q_row_update, q_row_delete)
	stmt, err := dbx.Prepare("insert into my_active(addtime, hostname, port, sselect, uupdate, ddelete, iinsert, ccall, q_row_insert, q_row_update, q_row_delete) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	stmt.Exec(addtime, hostname, port, sselect, uupdate, ddelete, iinsert, ccall, q_row_insert, q_row_update, q_row_delete)

	//innodb io status
	i_data_read := (tmp2["Innodb_data_read"] - tmp1["Innodb_data_read"])/t
	i_data_reads :=  (tmp2["Innodb_data_reads"] - tmp1["Innodb_data_reads"])/t
	i_data_written := (tmp2["Innodb_data_written"] - tmp1["Innodb_data_wirtten"])/t
	i_data_writes  := (tmp2["Innodb_data_writes"] - tmp1["Innodb_data_wirtes"])/t
	i_page_read := (tmp2["Innodb_pages_read"] - tmp1["Innodb_pages_read"])/t
	i_page_written := (tmp2["Innodb_pages_written"] - tmp1["Innodb_pages_written"])/t
	i_page_created := (tmp2["Innodb_pages_created"] - tmp1["Innodb_pages_created"])/t
	i_bp_read_requests := (tmp2["Innodb_buffer_pool_read_requests"] - tmp1["Innodb_buffer_pool_read_requests"])/t
	i_bp_reads := (tmp2["Innodb_buffer_pool_reads"] - tmp1["Innodb_buffer_pool_reads"])/t
	i_log_written := (tmp2["Innodb_os_log_written"] - tmp1["Innodb_os_log_written"])/t
	
	//innodb row active
	i_row_read := (tmp2["Innodb_rows_read"] - tmp1["Innodb_rows_read"])/t
	i_row_insert := (tmp2["Innodb_rows_inserted"] - tmp1["Innodb_rows_inserted"])/t
	i_row_update := (tmp2["Innodb_rows_updated"] - tmp1["Innodb_rows_updated"])/t
	i_row_delete := (tmp2["Innodb_rows_deleted"] - tmp1["Innodb_rows_deleted"])/t
	i_bp_hit := 0
	if ( i_bp_read_requests == 0  || i_bp_reads == 0 ){
		i_bp_hit = 100
	}else{
		i_bp_hit = (100  - ( 1000*(tmp2["Innodb_buffer_pool_reads"] - tmp1["Innodb_buffer_pool_reads"])/(tmp2["Innodb_buffer_pool_read_requests"] - tmp1["Innodb_buffer_pool_read_requests"])/10 ))
	}
	
	i_file := "./output/innodb_status.csv"
	io_fd, er2 :=os.OpenFile(i_file,os.O_RDWR|os.O_APPEND, 0660)
	defer io_fd.Close()
	if er2 != nil{
		io_fd, _ = os.Create(i_file)
	}
	fmt.Fprintf(io_fd, "%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n",addtime, i_bp_hit, i_data_read, i_data_reads, i_data_written, i_data_writes, i_page_read, i_page_written, i_page_created, i_bp_read_requests, i_bp_reads, i_log_written, i_row_read, i_row_insert, i_row_update, i_row_delete)
        i_stmt, err := dbx.Prepare("insert into my_innodb_status(addtime, hostname, port, i_bp_hit, i_data_read, i_data_reads, i_data_written, i_data_writes, i_page_read, i_page_written, i_page_created, i_bp_read_requests, i_bp_reads,i_log_written, i_row_read, i_row_insert, i_row_update, i_row_delete) values( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer i_stmt.Close()
	i_stmt.Exec(addtime, hostname, port, i_bp_hit, i_data_read, i_data_reads, i_data_written, i_data_writes, i_page_read, i_page_written, i_page_created, i_bp_read_requests, i_bp_reads, i_log_written, i_row_read, i_row_insert, i_row_update, i_row_delete)

	//thread
	t_run := tmp2["Threads_running"]
	t_cached := tmp2["Threads_cached"]
	t_conned := tmp2["Threads_connected"]
	t_max_conn := tmp2["Max_used_connections"]
	t_aborted_conn := (tmp2["Aborted_connects"] - tmp1["Aborted_connects"])/t
	t_thread_created := (tmp2["Threads_created"] - tmp1["Threads_created"])/t
	
	t_file := "./output/thread_conn.csv"
	t_fd, er3 := os.OpenFile(t_file, os.O_RDWR|os.O_APPEND, 0660)
	defer t_fd.Close()
	if  er3 != nil{
		t_fd, _ = os.Create(t_file)
	}
	fmt.Fprintf(t_fd, "%d,%d,%d,%d,%d,%d,%d\n", addtime, t_run, t_conned, t_cached, t_max_conn, t_thread_created, t_aborted_conn)
	t_stmt, err := dbx.Prepare("insert into my_thread(addtime, hostname, port, t_run, t_conned, t_cached, t_max_conn, t_thread_created, t_aborted_conn) values(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t_stmt.Exec(addtime, hostname, port, t_run, t_conned, t_cached, t_max_conn, t_thread_created, t_aborted_conn)

}


func OutputMySQL(user,password, port string , addtime int , mdb *comm.DBconn){
	_, ok := tmp1["Uptime"]
	if  ok != true {
		tmp1 = get_mystr(user,password,port)
	}
	time.Sleep(1*time.Second)
	tmp2 := get_mystr(user,password,port)

	my_active(tmp1, tmp2, addtime,port, mdb)
	tmp1 = tmp2
}
