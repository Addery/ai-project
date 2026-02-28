package main

import (
	"encoding/json"
	"fmt"
)

/*
	json tag
		这个不写json标签转换为json的话，字段名就是属性的名字
		小写的属性也不会转换
	空
		如果再转json的时候，我不希望某个字段被转出来，我可以写一个 -
	omitempty
		空值省略
*/

type Peasant struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
}

func main() {
	p := Peasant{
		Name:     "Addery",
		Age:      0,
		Password: "123456",
	}
	byteData, _ := json.Marshal(p)
	fmt.Println(string(byteData)) //{"name":"Addery"}
}
