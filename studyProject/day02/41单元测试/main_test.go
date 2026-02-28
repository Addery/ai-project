package main

import (
	"fmt"
	"os"
	"testing"
)

// 测试前执行
func setup() {
	fmt.Println("Before all tests")
}

// 测试后执行
func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

// 必须叫这个名字  测试主入口
func TestMain(m *testing.M) {
	// 测试前执行
	setup()
	code := m.Run()
	// 测试后执行
	teardown()
	os.Exit(code)
}
