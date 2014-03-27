package sys

import (
	"bytes"
	"fmt"
	"io"
	//	"time"
	"comm"
	"os"
	"os/exec"
	"strconv"
	"strings"
	//	"encoding/csv"
)

func get_diskispnfo(mdev string) (tmp []string) {
	cmd1 := exec.Command("bash", "-c", "df -Pm | grep ^/ | grep -v proc | grep -v none")
	cmd2 := exec.Command("grep", mdev)
	r, w := io.Pipe()
	cmd1.Stdout = w
	cmd2.Stdin = r

	var out bytes.Buffer
	cmd2.Stdout = &out
	cmd1.Start()
	cmd2.Start()
	cmd1.Wait()
	w.Close()
	cmd2.Wait()
	tmp = strings.Fields(out.String())
	return
}

func OutputDisksp(t int, mnode string, mdb *comm.DBconn) {
	diskspfile := "./output/diskspinfo.cvs"
	sp_fd, err := os.OpenFile(diskspfile, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		sp_fd, err = os.Create(diskspfile)
	}
	defer sp_fd.Close()
	nsp := get_diskispnfo(mnode)
	mdev := strings.Split(nsp[0], "/")[2]
	m_total, _ := strconv.Atoi(nsp[1])
	m_use, _ := strconv.Atoi(nsp[2])
	m_avail, _ := strconv.Atoi(nsp[3])
	m_use_pct, _ := strconv.Atoi(strings.TrimSuffix(nsp[4], "%"))
	m_node := nsp[5]
	fmt.Fprintf(sp_fd, "%d,%s,%s,%d,%d,%d,%d\n", t, m_node, mdev, m_total, m_use, m_avail, m_use_pct)

	hostname, _ := os.Hostname()

	dbx := mdb.GetMySQLconn()
	defer mdb.PutMySQL(dbx)

	stmt, err := dbx.Prepare("insert into sys_diskspinfo(addtime, hostname, mount_node, mdev, node_size, node_use, node_avail, node_use_pct) values(?, ?, ?, ?, ?, ?, ?, ?) ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	stmt.Exec(t, hostname, m_node, mdev, m_total, m_use, m_avail, m_use_pct)
}
