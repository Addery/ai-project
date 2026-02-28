package main

import (
	"fmt"
	"studyProject/day04_gorm/comm"
	"studyProject/day04_gorm/entity"
)

func main() {
	comm.Connect()
	db := comm.DB

	err := db.AutoMigrate(&entity.UserModel{})
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	fmt.Println("表结构迁移成功")
}
