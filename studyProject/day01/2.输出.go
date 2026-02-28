package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "你好")          // 可以作为任何值的占位符输出
	fmt.Printf("%v %T\n", "枫枫", "枫枫") // 打印类型
	fmt.Printf("%d\n", 3)             // 整数
	fmt.Printf("%.2f\n", 1.25)        // 小数
	fmt.Printf("%s\n", "哈哈哈")         // 字符串
	fmt.Printf("%#v\n", "")           // 用go的语法格式输出，很适合打印空字符串

	// 将格式化之后的内容赋值给一个变量
	var name = fmt.Sprintf("%v", "你好")
	fmt.Println(name)
}
