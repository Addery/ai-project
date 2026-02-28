package main

import (
	"fmt"
	"sync"
)

var num int
var wait sync.WaitGroup

func add() {
	for i := 0; i < 1000000; i++ {
		num++
	}
	wait.Done()
}
func reduce() {
	for i := 0; i < 1000000; i++ {
		num--
	}
	wait.Done()
}

func main() {
	wait.Add(2)
	go add()
	go reduce()
	wait.Wait()
	fmt.Println(num)

}
