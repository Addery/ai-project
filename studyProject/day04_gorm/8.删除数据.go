package main

import (
	"fmt"
	"studyProject/day04_gorm/comm"
	"studyProject/day04_gorm/entity"

	"gorm.io/gorm"
)

func deleteMethod(db *gorm.DB) {
	//var user = entity.UserModel{ID: 7}
	//db.Delete(&user)

	//db.Delete(&entity.UserModel{ID: 6})

	db.Delete(&entity.UserModel{}, []int{2, 3})
}

func softDeleteMethod(db *gorm.DB) {
	//db.Delete(&entity.UserModel{ID: 16})

	// 直接查找找不到软删除的记录，查看软删除记录方法如下
	var users []entity.UserModel
	db.Unscoped().Find(&users)
	fmt.Println(users)

	// 硬删除
	db.Unscoped().Delete(&entity.UserModel{ID: 15})
}

func main() {
	comm.Connect()
	//deleteMethod(comm.DB)
	softDeleteMethod(comm.DB)
}
