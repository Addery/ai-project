package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.0f\n", math.Pow(2, 63))
	var n1 int = 9223372036854775807
	fmt.Println(n1)
	//var n2 int = 9223372036854775808 // 看它报不报错
	//fmt.Println(n2)

	var c1 = 'a'
	var c2 = 97
	fmt.Println(c1) // 直接打印都是数字
	fmt.Println(c2)

	fmt.Printf("%c %d\n", c1, c2) // 以字符的格式打印

	var r1 rune = '中'
	fmt.Printf("%c\n", r1)

	// 常用转义字符
	fmt.Println("枫枫\t知道")              // 制表符
	fmt.Println("枫枫\n知道")              // 回车
	fmt.Println("\"枫枫\"知道")            // 双引号
	fmt.Println("枫枫\r知道")              // 回到行首
	fmt.Println("C:\\pprof\\main.exe") // 反斜杠

	// 多行字符串，在``这个里面，再出现转义字符就会原样输出了
	var s = `今天\n
天气
真好
`
	fmt.Println(s)

	// 0值
	var a1 int
	var a2 float32
	var a3 string
	var a4 bool

	fmt.Printf("%#v\n", a1)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)
}
