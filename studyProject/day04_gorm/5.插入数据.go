package main

import (
	"fmt"
	"studyProject/day04_gorm/comm"
	"studyProject/day04_gorm/entity"
)

func InsertData() {
	user := entity.UserModel{
		Name: "张三",
	}
	comm.Connect()
	fmt.Println("Create方法执行前")
	err := comm.DB.Create(&user).Error
	fmt.Println("Create方法执行后")
	fmt.Println(user, err)
}

func InsertBatchData() {
	var users = []entity.UserModel{
		{
			Name: "小刘",
		},
		{
			Name: "小白",
		},
	}
	comm.Connect()
	err := comm.DB.Create(&users).Error
	fmt.Println(users, err)
}

func main() {
	//InsertData()
	InsertBatchData()
}
