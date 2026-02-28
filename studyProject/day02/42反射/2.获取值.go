package main

import (
	"fmt"
	"reflect"
)

func refValue(obj any) {
	value := reflect.ValueOf(obj)
	fmt.Println(value, value.Type())
	// 去判断具体的类型
	switch value.Kind() {
	case reflect.Int:
		fmt.Println(value.Int())
	case reflect.Struct:
		fmt.Println(value.Interface())
	case reflect.String:
		fmt.Println(value.String())
	}
}

func main() {
	refValue(struct {
		Name string
	}{Name: "枫枫"})

	name := "枫枫"
	refValue(name)

	refValue(1)
}
