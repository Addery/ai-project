package main

import "fmt"

//const (
//	SuccessCode = 0
//	ErrorCode   = 1
//)
//
//func getMsg(code int) (msg string) {
//	switch code {
//	case SuccessCode:
//		msg = "success"
//	case ErrorCode:
//		msg = "error"
//	default:
//		msg = "error"
//	}
//	return
//}
//
//func Server(str string) (code int, msg string) {
//	if str == "1" {
//		return SuccessCode, getMsg(SuccessCode)
//	}
//	if str == "2" {
//		return ErrorCode, getMsg(ErrorCode)
//	}
//	return ErrorCode, getMsg(ErrorCode)
//}

type Code int

const (
	SuccessCode Code = 0
	ErrorCode   Code = 1
)

func (c Code) getMsg() (msg string) {
	switch c {
	case SuccessCode:
		msg = "success"
	case ErrorCode:
		msg = "error"
	default:
		msg = "error"
	}
	return
}

func (c Code) getCodeAndMsg() (code Code, msg string) {
	return c, c.getMsg()
}

func Server(str string) (code Code, msg string) {
	if str == "1" {
		return code.getCodeAndMsg()
	}
	if str == "2" {
		return code.getCodeAndMsg()
	}
	return code.getCodeAndMsg()
}

func main() {
	fmt.Println(Server("1"))
}
