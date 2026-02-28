package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func event() {
	defer close(done)
	fmt.Println("event执行中")
	time.Sleep(2 * time.Second)
	fmt.Println("event执行结束")
}

func main() {
	go event()

	select {
	case <-done:
		fmt.Println("协程执行结束")
	case <-time.After(1 * time.Second):
		fmt.Println("协程超时")
		return
	}
}
