package main

import (
	"fmt"
	"reflect"
)

func refSetValue(obj any) {
	value := reflect.ValueOf(obj)
	elem := value.Elem()
	switch elem.Kind() {
	case reflect.String:
		elem.SetString("Hello")
	}
}

func main() {
	msg := "test.txt"
	refSetValue(&msg)
	fmt.Println(msg)
}
