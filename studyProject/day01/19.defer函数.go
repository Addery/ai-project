package main

import "fmt"

func Func() {
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
}

/*
	defer函数：
		关键字 defer 用于注册延迟调用
		这些调用直到 return 前才被执。因此，可以用来做资源清理
		多个defer语句，按先进后出的方式执行，谁离return近谁先执行
		defer语句中的变量，在defer声明时就决定了
*/

func main() {
	defer fmt.Println("defer4")
	Func()
	defer fmt.Println("defer3")
}
