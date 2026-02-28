package main

import "fmt"

/*
	和自定义类型很像，但是有一些地方和自定义类型有很大差异
		不能绑定方法
		打印类型还是原始类型
		和原始类型比较，类型别名不用转换
*/

type AliasCode = int
type MyCode int

const (
	TestCode      MyCode    = 0
	TestAliasCode AliasCode = 0
)

// MyCodeMethod 自定义类型可以绑定自定义方法
func (m MyCode) MyCodeMethod() {

}

// MyAliasCodeMethod 类型别名 不可以绑定方法
//func (m AliasCode) MyAliasCodeMethod() {
//
//}

func main() {
	// 类型别名，打印它的类型还是原始类型
	fmt.Printf("%T %T \n", TestCode, TestAliasCode) // main.MyCode int
	// 可以直接和原始类型比较
	var i int
	fmt.Println(TestAliasCode == i)
	fmt.Println(int(TestCode) == i) // 必须转换之后才能和原始类型比较
}
