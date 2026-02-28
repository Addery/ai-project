package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 实例化一个http客户端
	client := new(http.Client)
	// 构造请求对象
	req, _ := http.NewRequest("GET", "http://localhost:8080/index", nil)
	// 发请求
	res, _ := client.Do(req)
	// 获取响应
	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))

}
