package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNum(n, f int)(res []int) {
	
	intRd:=0  
	count := 0

	for i := count; i < n; {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		intRd = r.Intn(f) + 1
		flag := 0
		for _,v:=range res {
			if v == intRd {
				flag = 1
				break
			}
		}
		if flag == 0 {
			res=append(res,  intRd)
			count++
                        i++
		}
	}
	return 

}

func main() {
	a := randomNum(100,4999 )
	fmt.Println(a)
}
