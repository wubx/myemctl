package comm

import (
	"strconv"
)

func Str_trade(tmp2, tmp1 string) (r int) {
	t1, _ := strconv.Atoi(tmp1)
	t2, _ := strconv.Atoi(tmp2)
	r = t2 - t1
	return
}
