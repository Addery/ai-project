package main

import "fmt"

func main() {
	way1()
	way2()
	way3()
}

func way1() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age <= 35 {
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")

}

func way2() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	if age <= 18 {
		if age <= 0 {
			fmt.Println("未出生")
		} else {
			fmt.Println("未成年")
		}
	} else {
		if age <= 35 {
			fmt.Println("青年")
		} else {
			fmt.Println("中年")
		}
	}
}
func way3() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	if age <= 0 {
		fmt.Println("未出生")
	}
	if age > 0 && age <= 18 {
		fmt.Println("未成年")
	}
	if age > 18 && age <= 35 {
		fmt.Println("青年")
	}
	if age > 35 {
		fmt.Println("中年")
	}
}
