package comm
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var MAX_POOL_SIZE=10
var MySQLPool chan *sql.DB

type DBconn struct {
	User string
	Password string
	Hostip string
	Port string
	DBname string
}
func (db DBconn )GetMySQLconn() (*sql.DB){
	if MySQLPool == nil {
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}

	if ( len(MySQLPool) == 0){
		go func(){
			for i:= 0; i < MAX_POOL_SIZE/2; i++{
				mysqlconn, err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",db.User,db.Password, db.Hostip,db.Port,db.DBname))
				if err != nil{
					panic(err)
				}
				db.PutMySQL(mysqlconn)
			}
		}()
	}
	return <- MySQLPool
}

func (db DBconn) PutMySQL(conn *sql.DB){
	if MySQLPool == nil{
		MySQLPool = make(chan *sql.DB, MAX_POOL_SIZE)
	}

	if (len(MySQLPool) == MAX_POOL_SIZE){
		conn.Close()
		return
	}
	MySQLPool <- conn
}
