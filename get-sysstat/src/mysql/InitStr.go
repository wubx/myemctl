package mysql

import (
	"comm"
	"fmt"
)

func InitStr(mdb *comm.DBconn) {
	dbx := mdb.GetMySQLconn()
	defer mdb.PutMySQL(dbx)

	cpu_str := `
	CREATE TABLE IF NOT EXISTS sys_cpuinfo (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		addtime INT UNSIGNED NOT NULL,
		hostname VARCHAR(45) NOT NULL DEFAULT 'local',
		cpu_user_pct SMALLINT UNSIGNED NOT NULL DEFAULT 0,
		cpu_nice_pct SMALLINT UNSIGNED NOT NULL DEFAULT 0,
		cpu_sys_pct SMALLINT NOT NULL DEFAULT 0,
		cpu_idle_pct SMALLINT NOT NULL DEFAULT 0,
		cpu_iowait_pct SMALLINT UNSIGNED NOT NULL DEFAULT 0,
		cpu_irq_pct SMALLINT UNSIGNED NOT NULL DEFAULT 0,
		cpu_softirq_pct SMALLINT NOT NULL DEFAULT 0,
		cpu_all_irq_pct SMALLINT UNSIGNED NOT NULL DEFAULT 0,
		cpu_total_pct SMALLINT UNSIGNED NOT NULL DEFAULT 0,
		PRIMARY KEY (id),
		INDEX idx_addtime_hostname (addtime , hostname))
		`
	_, err := dbx.Query(cpu_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}
	load_str := `
	CREATE TABLE IF NOT EXISTS sys_load (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		addtime INT UNSIGNED NOT NULL,
		hostname VARCHAR(45) NULL,
		la1 DECIMAL(5,2) NULL,
		la5 DECIMAL(5,2) NULL,
		la15 DECIMAL(5,2) NULL,
		pid_and_num VARCHAR(30) NULL,
		run_pid VARCHAR(30) NULL,
		PRIMARY KEY (id),
		INDEX idx_addtime_hostname (addtime, hostname))
		`

	_, err = dbx.Query(load_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}
	diskio_str := `
	CREATE TABLE IF NOT EXISTS sys_diskioinfo (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		addtime INT NOT NULL,
		hostname VARCHAR(45) NOT NULL DEFAULT 'local',
		dev_name VARCHAR(45) NOT NULL,
		dev_io_pers INT UNSIGNED NOT NULL DEFAULT 0,
		dev_read_pers INT UNSIGNED NULL DEFAULT 0,
		dev_readmerge_pers INT UNSIGNED NULL,
		dev_readsector_pers INT UNSIGNED NULL,
		dev_read_use_pers INT UNSIGNED NULL,
		dev_write_pers INT UNSIGNED NULL,
		dev_writemerge_pers INT UNSIGNED NULL,
		dev_writesector_pers INT NULL,
		dev_write_use_pers INT UNSIGNED NULL,
		PRIMARY KEY (id),
		INDEX idx_addtime_hostname (addtime, hostname))
		`
	_, err = dbx.Query(diskio_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}

	disksp_str := `
		CREATE TABLE IF NOT EXISTS sys_diskspinfo (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			addtime INT UNSIGNED NULL,
			hostname VARCHAR(45) NULL,
			mount_node VARCHAR(45) NULL,
			mdev VARCHAR(45) NULL,
			node_size INT UNSIGNED NULL,
			node_use INT UNSIGNED NULL,
			node_avail INT UNSIGNED NULL,
			node_use_pct INT UNSIGNED NULL,
			PRIMARY KEY (id),
			INDEX idx_addtime_hostname (addtime, hostname))
			`
	_, err = dbx.Query(disksp_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}

	mem_str := `
		CREATE TABLE IF NOT EXISTS sys_mem (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			addtime INT UNSIGNED NULL,
			hostname VARCHAR(45) NULL,
			mem_total INT NULL,
			mem_free INT NULL,
			PRIMARY KEY (id),
			INDEX idx_addtime_hostname (addtime, hostname))
			`
	_, err = dbx.Query(mem_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}

	my_avtive_str := `
		CREATE TABLE IF NOT EXISTS my_active (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			addtime INT NOT NULL,
			hostname VARCHAR(45) NOT NULL DEFAULT 'local',
			port INT NOT NULL DEFAULT 3306,
			sselect INT NULL,
			uupdate INT NULL,
			ddelete INT NULL,
			iinsert INT NULL,
			ccall INT NULL,
			q_row_insert INT NULL,
			q_row_update INT NULL,
			q_row_delete INT NULL,
			PRIMARY KEY (id),
			INDEX idx_addtime_hostname_port (addtime, hostname, port))
			`
	_, err = dbx.Query(my_avtive_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}

	my_innodb_str := `
		CREATE TABLE IF NOT EXISTS my_innodb_status (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			addtime INT UNSIGNED NOT NULL,
			hostname VARCHAR(45) NOT NULL DEFAULT 'local',
			port INT NOT NULL DEFAULT 3306,
			i_bp_hit INT NULL,
			i_data_read INT NULL,
			i_data_reads INT NULL,
			i_data_written INT NULL,
			i_data_writes INT NULL,
			i_page_read INT NULL,
			i_page_written INT NULL,
			i_page_created INT NULL,
			i_bp_read_requests INT NULL,
			i_bp_reads INT NULL,
			i_log_written INT NULL,
			i_row_read INT NULL,
			i_row_insert INT NULL,
			i_row_update INT NULL,
			i_row_delete INT NULL,
			PRIMARY KEY (id),
			INDEX idx_addtime_hostname_port (addtime, hostname, port))
			`
	_, err = dbx.Query(my_innodb_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}

	my_thread_str := `
		CREATE TABLE IF NOT EXISTS my_thread (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			addtime INT UNSIGNED NOT NULL,
			hostname VARCHAR(45) NOT NULL DEFAULT 'local',
			port INT NOT NULL DEFAULT 3306,
			t_run INT NULL,
			t_conned INT NULL,
			t_cached INT NULL,
			t_max_conn INT NULL,
			t_thread_created INT NULL,
			t_aborted_conn INT NULL,
			PRIMARY KEY (id),
			INDEX idx_addtime_hostname_port (addtime, hostname, port))
			`
	_, err = dbx.Query(my_thread_str)
	if err != nil {
		fmt.Println("Eorr:", err)
		return
	}

	fmt.Println("Create OK")
}
