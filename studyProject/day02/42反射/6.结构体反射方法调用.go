package main

import (
	"fmt"
	"reflect"
)

type User struct {
}

func (User) Call(name string) {
	fmt.Println("call方法执行了", name)
}

func Call(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()

	for i := 0; i < v.NumMethod(); i++ {
		method := t.Method(i)
		if method.Name != "Call" {
			continue
		}
		v.Method(i).Call([]reflect.Value{
			reflect.ValueOf("addery"),
		})
	}
}

func main() {
	u := User{}
	Call(&u)
}
