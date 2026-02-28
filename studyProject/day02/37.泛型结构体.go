package main

import (
	"encoding/json"
	"fmt"
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main() {
	type User struct {
		Name string `json:"name"`
	}

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	//user := Response{
	//  Code: 0,
	//  Msg:  "成功",
	//  Data: User{
	//    Name: "枫枫",
	//  },
	//}
	//byteData, _ := json.Marshal(user)
	//fmt.Println(string(byteData))
	//userInfo := Response{
	//  Code: 0,
	//  Msg:  "成功",
	//  Data: UserInfo{
	//    Name: "枫枫",
	//    Age:  24,
	//  },
	//}
	//byteData, _ = json.Marshal(userInfo)
	//fmt.Println(string(byteData))

	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)
	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫","age":24}}`), &userInfoResponse)
	fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)
}
