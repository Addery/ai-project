package main

import (
	"fmt"
	"sort"
)

/*
go里面的数组，长度被限制死了，所以不经常用
所以go出了一个数组plus，叫做slice（切片）
切片（Slice）相较于数组更灵活，因为在声明切片后其长度是可变的
*/
func main() {
	var list []string

	fmt.Println(list == nil) // true

	list = append(list, "hello")
	list = append(list, "world")
	fmt.Println(list)
	fmt.Println(len(list))
	list[1] = "hello"
	fmt.Println(list)

	// 定义一个字符串切片 make([]type, length, capacity)
	var list1 = make([]string, 0)
	fmt.Println(list1, len(list1), cap(list1))
	fmt.Println(list1 == nil) // false

	list2 := make([]int, 2, 2)
	fmt.Println(list2, len(list2), cap(list2))

	// 切片排序
	var list3 = []int{4, 5, 3, 2, 7}
	fmt.Println("排序前：", list3)

	sort.Ints(list3)
	fmt.Println("升序：", list3)

	sort.Sort(sort.Reverse(sort.IntSlice(list3)))
	fmt.Println("降序：", list3)
}
