package main

import "fmt"

/*
	Go语言中的map(映射、字典)是一种内置的数据结构，它是一个无序的key-value对的集合
	map的key必须是基本数据类型，value可以是任意类型
*/

func main() {
	var m1 map[string]string

	//m1 = make(map[string]string)
	m1 = map[string]string{}

	m1["name"] = "zzy"
	fmt.Println(m1)

	fmt.Println(m1["name"])

	delete(m1, "name")
	fmt.Println(m1)

	var m2 = make(map[string]string)
	var m3 = map[string]string{}
	fmt.Println(m2)
	fmt.Println(m3)

	// 声明并赋值
	var m4 = map[string]int{
		"age": 21,
	}
	var age1 = m4["age1"]
	fmt.Println(age1) // 0
	var age2, ok2 = m4["age2"]
	fmt.Println(age2, ok2) // 0 false
	var age3, ok3 = m4["age"]
	fmt.Println(age3, ok3) // 21 true
}
