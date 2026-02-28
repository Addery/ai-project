package main

import (
	"fmt"
	"sync"
)

var num1 int
var wait1 sync.WaitGroup
var lock sync.Mutex

func add1() {
	// 谁先抢到了这把锁，谁就把它锁上，一旦锁上，其他的线程就只能等着
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num1++
	}
	lock.Unlock()
	wait1.Done()
}
func reduce1() {
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num1--
	}
	lock.Unlock()
	wait1.Done()
}

func main() {
	wait1.Add(2)
	go add1()
	go reduce1()
	wait1.Wait()
	fmt.Println(num1)

}
