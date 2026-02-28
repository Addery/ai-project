package main

import (
	"fmt"
	"time"
)

func main() {
	//for1()
	//for2()
	//for3()
	//for4()
	//for5()
	//for6()
	//for7()
	for8()
}

func for1() {
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 每隔1秒打印当前的时间
func for2() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 年月日时分秒的固定格式
	}
}

// 由于golang没有while循环，如果需要，则是由for循环稍微变化得来
func for3() {
	i := 0
	sum := 0
	for i < 100 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

// do-while模式就是先执行一次循环体，再判断
func for4() {
	i := 0
	sum := 0
	for {
		sum += i
		i++
		if i > 100 {
			break
		}
	}
	fmt.Println(sum)
}

// 遍历切片
func for5() {
	list := []string{"a", "b", "c", "d"}
	for index, s := range list {
		fmt.Println(index, s)
	}
}

// 遍历map
func for6() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func for7() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * % d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func for8() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				// 去掉 列比行大的数据
				continue
			}
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
