package main

import (
	"errors"
	"fmt"
)

// 无返回值
func fun1() {
	return // 也可以不写
}

// 单返回值
func fun2() int {
	return 1
}

// 多返回值
func fun3() (int, error) {
	return 0, errors.New("错误")
}

// 命名返回值
func fun4() (res string) {
	//res = "hello"
	return // 相当于先定义再赋值
	//return "abc"
}

func main() {
	fmt.Println(fun4())
}
