package main

import (
	"fmt"
	"studyProject/day04_gorm/comm"
	"studyProject/day04_gorm/entity"

	"gorm.io/gorm"
)

func selectAllUser(db *gorm.DB) {
	var users []entity.UserModel
	db.Find(&users)
	fmt.Println(users)
}

func selectUserByCondition(db *gorm.DB) {
	var users []entity.UserModel
	db.Find(&users, "name = ?", "张三")
	fmt.Println(users)
}

func selectOnlyUser(db *gorm.DB) {
	var user entity.UserModel
	db.Take(&user)
	//db.First(&user)
	//db.Last(&user)

	fmt.Println(user)
}

func selectUserError(db *gorm.DB) {
	user := entity.UserModel{}
	// 查不到会报错的写法
	//err := db.Take(&user, "name = ?", "stone").Error
	//if err == gorm.ErrRecordNotFound {
	//	fmt.Println("不存在的记录")
	//}

	// 不会报错的写法
	err := db.Limit(1).Find(&user, "name = ?", "stone").Error
	fmt.Println(err)
}

/*
打印实际的sql
*/
func printRealSQL(db *gorm.DB) {
	var user entity.UserModel
	db.Debug().Take(&user, "id = ?", 1)
	fmt.Println(user)
}

func main() {
	comm.Connect()
	//selectAllUser(comm.DB)
	//selectUserByCondition(comm.DB)
	//selectOnlyUser(comm.DB)
	//selectUserError(comm.DB)
	printRealSQL(comm.DB)
}
