package main

import (
	"errors"
	"fmt"
)

func Parent() error {
	err := method() // 遇到错误向上抛
	return err
}
func method() error {
	return errors.New("出错了")
}

func main() {
	fmt.Println(Parent())
}
