// package main
//
// import "fmt"
//
//	func main() {
//		fmt.Println("请输入要执行的操作：")
//		fmt.Println(`1：登录
//
// 2：个人中心
// 3：注销`)
//
//		var num int
//		fmt.Scan(&num)
//		var funcMap = map[int]func(){
//			1: func() {
//				fmt.Println("登录")
//			},
//			2: func() {
//				fmt.Println("个人中心")
//			},
//			3: func() {
//				fmt.Println("注销")
//			},
//		}
//		funcMap[num]()
//	}
package main

import "fmt"

func login() {
	fmt.Println("登录")
}

func userCenter() {
	fmt.Println("个人中心")
}

func logout() {
	fmt.Println("注销")
}

func main() {
	fmt.Println("请输入要执行的操作：")
	fmt.Println(`1：登录
2：个人中心
3：注销`)
	var num int
	fmt.Scan(&num)
	var funcMap = map[int]func(){
		1: login,
		2: userCenter,
		3: logout,
	}
	funcMap[num]()
}
