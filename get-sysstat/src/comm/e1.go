package  comm
import (
	"fmt"
	"time"
)
var  t1  int64
func init(){
	t1 =  time.Now().Unix()
}
func E1(){
	fmt.Println("Hello e1", t1)
	time.Sleep(1*time.Second)
	t2 := time.Now().Unix()
	fmt.Println("New time: ",t2)
	t1  =  t2
}
