package main

import "fmt"

/*
init()函数是一个特殊的函数，存在以下特性：

	不能被其他函数调用，而是在main函数执行之前，自动被调用
	init函数不能作为参数传入
	不能有传入参数和返回值
	一个go文件可以有多个init函数，谁在前面谁就先执行
*/
func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init3")
}

func main() {
	fmt.Println("main")
}
