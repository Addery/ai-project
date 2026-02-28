package main

import (
	"fmt"
	"studyProject/day01/version"
)

var global = "全局变量，可以不使用"

func main() {

	// 定义
	var name string
	// 赋值
	name = "分布赋值"
	// 使用
	fmt.Println(name)

	var userName string = "直接赋值"
	fmt.Println(userName)

	var userName1 = "省略类型"
	fmt.Println(userName1)

	userName2 := "简短声明"
	fmt.Println(userName2)

	var local = "局部变量，必须使用"
	fmt.Println(local)

	// 定义多个变量
	var name1, name2, name3 string
	fmt.Println(name1, name2, name3)

	var a1, a2 = "a1", "a2"
	fmt.Println(a1, a2)

	a3, a4 := "a3", "a4"
	fmt.Println(a3, a4)

	var (
		a5 string = "a5"
		a6        = "a6"
	)
	fmt.Println(a5, a6)

	// 定义常量
	const constVar string = "定义常量，不能再修改"
	fmt.Println(constVar)

	fmt.Println(version.Version)
	//fmt.Println(version.name)
}
