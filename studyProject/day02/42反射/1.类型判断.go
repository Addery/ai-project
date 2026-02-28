package main

import (
	"fmt"
	"reflect"
)

func refType(obj any) {
	typeObj := reflect.TypeOf(obj)
	fmt.Println(typeObj, typeObj.Kind())
	// 去判断具体的类型
	switch typeObj.Kind() {
	case reflect.Slice:
		fmt.Println("切片")
	case reflect.Map:
		fmt.Println("map")
	case reflect.Struct:
		fmt.Println("结构体")
	case reflect.String:
		fmt.Println("字符串")
	}
}

func main() {
	refType(struct {
		Name string
	}{Name: "枫枫"})

	name := "枫枫"
	refType(name)

	refType([]string{"枫枫"})
}
