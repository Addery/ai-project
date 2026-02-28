package main

import "fmt"

func value(num int) {
	fmt.Println(&num) // 可以看到，这个n的内存地址和外面num的内存地址是明显不一样的
	fmt.Println(num)
	num = 2 // 这里的修改不会影响外面的num
}

func cite(num *int) {
	fmt.Println(num) // 内存值是一样的
	*num = 2         // 这里的修改会影响外面的num
}

/*
&是取地址，*是解引用，去这个地址指向的值
*/
func main() {
	num := 3
	fmt.Println(&num)
	value(num)
	fmt.Println("值传递结果：", num)

	cite(&num)
	fmt.Println(&num)
	fmt.Println("引用传递结果：", num)
}
