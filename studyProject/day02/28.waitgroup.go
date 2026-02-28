package main

import (
	"fmt"
	"sync"
	"time"
)

//var wait = sync.WaitGroup{}

func sing1(wait *sync.WaitGroup) {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
	wait.Done()
}

func main() {
	start := time.Now()
	var wait sync.WaitGroup
	wait.Add(4)
	go sing1(&wait)
	go sing1(&wait)
	go sing1(&wait)
	go sing1(&wait)

	wait.Wait()
	fmt.Println("全部结束，耗时：", time.Since(start))
}
