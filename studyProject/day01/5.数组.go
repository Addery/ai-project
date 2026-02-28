package main

import "fmt"

func main() {
	var array [3]int = [3]int{1, 2, 3}
	fmt.Println(array)

	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)

	var array2 = [...]int{1, 2, 3}
	fmt.Println(array2)

	array[0] = 10
	fmt.Println(array)

	var array3 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 拿到array3中倒数第二个元素
	fmt.Println(array3[len(array3)-2])
}
