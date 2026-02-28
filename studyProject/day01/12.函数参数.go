package main

import "fmt"

func add(n1 int, n2 int) {
	fmt.Println(n1, n2)
}

// 参数类型一样，可以合并在一起
func add1(n1, n2 int) {
	fmt.Println(n1, n2)
}

// 多个参数
func add2(numList ...int) {
	fmt.Println(numList)
}

func main() {
	add(1, 2)
	add1(1, 2)
	add2(1, 2)
	add2(1, 2, 3, 4)
}
