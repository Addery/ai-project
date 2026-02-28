package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int) // 声明并初始化一个长度为0的信道

func pay(name string, money int, wait *sync.WaitGroup) {
	defer wait.Done()
	fmt.Printf("%s 模拟购物中\n", name)
	time.Sleep(1 * time.Second)
	moneyChan <- money
	fmt.Println("购物结束")
}

func main() {
	var wait sync.WaitGroup
	var moneyList []int
	startTime := time.Now()

	wait.Add(3)
	go pay("张三", 2, &wait)
	go pay("李四", 3, &wait)
	go pay("王五", 5, &wait)

	go func() {
		defer close(moneyChan)
		wait.Wait()
	}()

	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}

	//for {
	//	money, ok := <-moneyChan
	//	if !ok {
	//		break
	//	}
	//	moneyList = append(moneyList, int(money))
	//}

	fmt.Println("购物耗时：", time.Since(startTime))
	fmt.Println("消费记录：", moneyList)

}
