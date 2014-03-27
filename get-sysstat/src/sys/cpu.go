package sys

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"strconv"
	"time"
	"encoding/csv"
	"comm"
)
type Tmpcpu struct {
	cpu_user int
	cpu_nice int
	cpu_sys int
	cpu_idle int
	cpu_iowait int
	cpu_irq int
	cpu_softirq int
	cpu_total int
}

type Cpuuse struct {
	cpu_user_pct int
	cpu_nice_pct int
	cpu_sys_pct int
	cpu_idle_pct int
	cpu_iowait_pct int
	cpu_irq_pct int
	cpu_softirq_pct int
	cpu_all_irq_pct int
	cpu_total_pct int
}

func (tmpcpu1 Tmpcpu) Cpupct(tmpcpu2 Tmpcpu) (cpuuse Cpuuse){
	var  diff Tmpcpu
	diff.cpu_user =  tmpcpu2.cpu_user - tmpcpu1.cpu_user
	diff.cpu_nice =  tmpcpu2.cpu_nice - tmpcpu1.cpu_nice
	diff.cpu_idle =  tmpcpu2.cpu_idle - tmpcpu1.cpu_idle
	diff.cpu_sys =  tmpcpu2.cpu_sys - tmpcpu1.cpu_sys
	diff.cpu_iowait =  tmpcpu2.cpu_iowait - tmpcpu1.cpu_iowait
	diff.cpu_irq =  tmpcpu2.cpu_irq - tmpcpu1.cpu_irq
	diff.cpu_softirq =  tmpcpu2.cpu_softirq - tmpcpu1.cpu_softirq
	diff.cpu_total =  tmpcpu2.cpu_total - tmpcpu1.cpu_total

	cpuuse.cpu_user_pct = 100000*diff.cpu_user/diff.cpu_total/10
	cpuuse.cpu_idle_pct = 100000*diff.cpu_idle/diff.cpu_total/10
	cpuuse.cpu_sys_pct = 100000*diff.cpu_sys/diff.cpu_total/10
	cpuuse.cpu_nice_pct = 100000*diff.cpu_nice/diff.cpu_total/10
	cpuuse.cpu_iowait_pct = 100000*diff.cpu_iowait/diff.cpu_total/10
	cpuuse.cpu_irq_pct = 100000*diff.cpu_irq/diff.cpu_total/10
	cpuuse.cpu_softirq_pct = 100000*diff.cpu_softirq/diff.cpu_total/10
	cpuuse.cpu_all_irq_pct = 100000*(diff.cpu_softirq+diff.cpu_irq)/diff.cpu_total/10
	cpuuse.cpu_total_pct =  10000 - cpuuse.cpu_idle_pct
	return
}


func GetCpu() ( t Tmpcpu ){
	cmd := exec.Command("bash", "-c","grep -m1 '^cpu' /proc/stat")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	//fmt.Printf("%q\n",strings.Fields(out.String()))
	tmp1 := strings.Fields(out.String())
	var tp Tmpcpu
	tp.cpu_user, _ = strconv.Atoi(tmp1[1])
	tp.cpu_nice, _ = strconv.Atoi(tmp1[2])
	tp.cpu_sys, _ = strconv.Atoi(tmp1[3])
	tp.cpu_idle, _ = strconv.Atoi(tmp1[4])
	tp.cpu_iowait, _ = strconv.Atoi(tmp1[5])
	tp.cpu_irq, _ = strconv.Atoi(tmp1[6])
	tp.cpu_softirq, _ = strconv.Atoi(tmp1[7])
	tp.cpu_total = tp.cpu_user + tp.cpu_nice + tp.cpu_sys + tp.cpu_idle + tp.cpu_iowait + tp.cpu_irq + tp.cpu_softirq
	return tp
}

func OutputCpu(t int,  mdb *comm.DBconn){
	var  tmp1_cpu, tmp2_cpu Tmpcpu
	tmp1_cpu = GetCpu()
	/*
	start := time.Now()
	h, min, _ := start.Clock()
	y, m, d := start.Date()
	*/
	cpuFile :="./output/cpu.csv"
	cpu_fd, err :=  os.OpenFile(cpuFile,os.O_RDWR|os.O_APPEND, 0660)
	defer cpu_fd.Close()

	if err != nil {
		cpu_fd, err =  os.Create(cpuFile)
	}
	buf := new(bytes.Buffer)
	r2 := csv.NewWriter(buf)
	s := make([]string, 10)

	time.Sleep(5*time.Second)
	tmp2_cpu = GetCpu()
	var  cpu_pct Cpuuse
	cpu_pct =  tmp1_cpu.Cpupct(tmp2_cpu)
	tmp1_cpu = tmp2_cpu
	//fmt.Println("recived :", t)
	s[0] =  strconv.Itoa(t)
	s[1] =  strconv.Itoa(cpu_pct.cpu_user_pct)
	s[2] =  strconv.Itoa(cpu_pct.cpu_nice_pct)
	s[3] =  strconv.Itoa(cpu_pct.cpu_sys_pct)
	s[4] =  strconv.Itoa(cpu_pct.cpu_idle_pct)
	s[5] =  strconv.Itoa(cpu_pct.cpu_iowait_pct)
	s[6] =  strconv.Itoa(cpu_pct.cpu_irq_pct)
	s[7] =  strconv.Itoa(cpu_pct.cpu_softirq_pct)
	s[8] =  strconv.Itoa(cpu_pct.cpu_all_irq_pct)
	s[9] =  strconv.Itoa(cpu_pct.cpu_total_pct)
	r2.Write(s)
	r2.Flush()
	//fmt.Println(buf)
	cpu_fd.WriteString(buf.String())
	
	hostname , _ := os.Hostname()
	dbx := mdb.GetMySQLconn()
	defer mdb.PutMySQL(dbx)
	stmt, err := dbx.Prepare("insert into sys_cpuinfo(addtime, hostname,cpu_user_pct, cpu_nice_pct,cpu_sys_pct, cpu_idle_pct, cpu_iowait_pct, cpu_irq_pct,cpu_softirq_pct, cpu_all_irq_pct, cpu_total_pct) values(?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	stmt.Exec(t,hostname, cpu_pct.cpu_user_pct, cpu_pct.cpu_nice_pct, cpu_pct.cpu_sys_pct, cpu_pct.cpu_idle_pct, cpu_pct.cpu_iowait_pct, cpu_pct.cpu_irq_pct, cpu_pct.cpu_softirq_pct, cpu_pct.cpu_all_irq_pct, cpu_pct.cpu_total_pct)
}
