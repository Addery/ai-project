package main

import (
	"fmt"
	"time"
)

func singing() {
	fmt.Println("唱歌")
	time.Sleep(1 * time.Second)
	fmt.Println("唱歌结束")
}

func main() {
	go singing()
	go singing()
	go singing()
	go singing()
	time.Sleep(2 * time.Second)
}
