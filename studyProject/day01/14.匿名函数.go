package main

import "fmt"

func main() {
	var add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(42, 13))
}
