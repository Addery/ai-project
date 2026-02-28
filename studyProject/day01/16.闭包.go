package main

import (
	"fmt"
	"time"
)

func awaitAdd(t int) func(...int) int {
	time.Sleep(time.Duration(t) * time.Second)
	return func(nums ...int) (res int) {
		for _, num := range nums {
			res += num
		}
		return
	}
}

func main() {
	//t1 := time.Now()
	//fmt.Println(awaitAdd(2)(1, 2, 3))
	//subTime := time.Since(t1)
	//fmt.Println(subTime)

	f := awaitAdd(2)
	t1 := time.Now()
	fmt.Println(f(1, 2, 3))
	subTime := time.Since(t1)
	fmt.Println(subTime)
}
