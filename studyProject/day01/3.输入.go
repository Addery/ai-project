package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入：")
	var name string
	fmt.Scan(&name)
	fmt.Println("你的输入是：", name)

	fmt.Println("请输入整数：")
	var age int
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}
