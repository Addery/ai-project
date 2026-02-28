package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Get(response http.ResponseWriter, request *http.Request) {
	// 获取参数
	fmt.Println(request.URL.String())

	byteDate, _ := json.Marshal(Response{
		Code: 200,
		Msg:  "ok",
		Data: map[string]any{},
	})

	response.Write(byteDate)

}

func Post(response http.ResponseWriter, request *http.Request) {
	// 获取参数
	bodyByteDate, _ := io.ReadAll(request.Body)
	fmt.Printf("%#v\n", string(bodyByteDate))

	byteDate, _ := json.Marshal(Response{
		Code: 200,
		Msg:  "ok",
		Data: map[string]any{},
	})

	response.Write(byteDate)
}

func main() {
	http.HandleFunc("/get", Get)
	http.HandleFunc("/post", Post)

	fmt.Println("http server running: http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
