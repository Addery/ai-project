package main

import (
	"fmt"
	"studyProject/day04_gorm/comm"
	"studyProject/day04_gorm/entity"

	"gorm.io/gorm"
)

func saveMethod(db *gorm.DB) {
	// 创建
	//user := entity.UserModel{
	//	Name: "xiaoZhang",
	//}
	//db.Save(&user)
	//fmt.Println(user)

	// 更新
	user := entity.UserModel{
		ID:   6,
		Name: "xiaoWang",
	}
	db.Save(&user)
	fmt.Println(user)
}

func updateMethod(db *gorm.DB) {
	// 走了BeforeUpdate
	db.Model(&entity.UserModel{}).
		Where("id = ?", 8).
		Update("name", "Addery")
}

func updateColumnMethod(db *gorm.DB) {
	// 不走BeforeUpdate
	db.Model(&entity.UserModel{}).
		Where("id = ?", 7).
		UpdateColumn("name", "小茗")
}

func updatesMethod(db *gorm.DB) {
	//var user = entity.UserModel{ID: 9}
	//db.Model(&user).Updates(entity.UserModel{
	//	Name: "小张",
	//})

	// 不更新零值
	//var user = entity.UserModel{ID: 9}
	//db.Model(&user).Updates(entity.UserModel{
	//	Name: "",
	//})

	// map会更新零值
	var user = entity.UserModel{ID: 9}
	db.Model(&user).Updates(map[string]any{
		"Name": "",
	})
}

/*
Expr通常用于获取原字段的数据
*/
func exprMethod(db *gorm.DB) {
	db.Model(&entity.UserModel{}).
		Where("id = ?", 7).
		UpdateColumn("age", gorm.Expr("age + ?", 1))
}

func main() {
	comm.Connect()
	exprMethod(comm.DB)
}
