package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int `json:"age"`
}

func parseStruct(obj any) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonField := field.Tag.Get("json")
		if jsonField == "" {
			jsonField = field.Name
		}
		fmt.Printf("Name: %s, type: %s, json: %s, value: %v\n", field.Name, field.Type, jsonField, v.Field(i))
	}
}

func main() {
	s := Student{
		Name: "test.txt",
		Age:  18,
	}
	parseStruct(s)
}
