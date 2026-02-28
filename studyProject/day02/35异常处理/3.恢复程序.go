package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 捕获异常，打印错误信息
			// 打印错误的堆栈信息
			fmt.Println(string(debug.Stack()))
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2]) // 肯定会有一个panic
}

func main() {

	read()

	fmt.Println("其他逻辑")
}
