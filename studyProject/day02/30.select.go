package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int) // 声明并初始化一个长度为0的信道
var nameChan = make(chan string)
var doneChan = make(chan struct{})

func send(name string, money int, wait *sync.WaitGroup) {
	defer wait.Done()
	fmt.Printf("%s 模拟购物中\n", name)
	time.Sleep(1 * time.Second)

	moneyChan1 <- money
	nameChan <- name

	fmt.Println("购物结束")
}

func main() {
	var wait sync.WaitGroup
	var moneyList []int
	var nameList []string
	startTime := time.Now()

	wait.Add(3)
	go send("张三", 2, &wait)
	go send("李四", 3, &wait)
	go send("王五", 5, &wait)

	go func() {
		defer close(doneChan)
		defer close(moneyChan1)
		defer close(nameChan)
		wait.Wait()
	}()

	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan:
				nameList = append(nameList, name)
			case <-doneChan:
				return
			}
		}
	}
	
	event()

	fmt.Println("购物耗时：", time.Since(startTime))
	fmt.Println("消费记录：", moneyList, nameList)

	//for {
	//	select {
	//	case money := <-moneyChan1:
	//		moneyList = append(moneyList, money)
	//	case name := <-nameChan:
	//		nameList = append(nameList, name)
	//	case <-doneChan:
	//		fmt.Println("购物耗时：", time.Since(startTime))
	//		fmt.Println("消费记录：", moneyList, nameList)
	//		return
	//	}
	//}

	//go func() {
	//	for money := range moneyChan1 {
	//		moneyList = append(moneyList, money)
	//	}
	//}()
	//
	//for name := range nameChan {
	//	nameList = append(nameList, name)
	//}

	//fmt.Println("购物耗时：", time.Since(startTime))
	//fmt.Println("消费记录：", moneyList, nameList)

}
