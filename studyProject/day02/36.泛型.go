package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// 泛型函数
func add2[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add2(1, 2))
}
