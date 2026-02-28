package main

import "fmt"

func main() {
	//switch1()
	//switch2()
	switch3()
}

func switch1() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age <= 18:
		fmt.Println("未成年")
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}
}

func switch2() {
	fmt.Println("请输入星期数字：")
	var week int
	fmt.Scan(&week)

	switch week {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	case 4:
		fmt.Println("周四")
	case 5:
		fmt.Println("周五")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("错误")
	}
}

func switch3() {
	/*
		我输入一个12，我希望它能输出满足的所有条件，例如我希望它输出，未成年，青年
	*/
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("未出生")
		fallthrough
	case age <= 18:
		fmt.Println("未成年")
		fallthrough
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}
}
