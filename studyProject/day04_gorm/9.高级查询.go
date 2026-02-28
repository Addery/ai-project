package main

import (
	"fmt"
	"studyProject/day04_gorm/comm"
	"studyProject/day04_gorm/entity"

	"gorm.io/gorm"
)

func whereExample(db *gorm.DB) {
	var users []entity.UserModel
	//db.Debug().Where("id = ?", 1).Take(&users)
	//fmt.Println(users)

	db.Where("id = 1 or age = 18").Find(&users)
	fmt.Println(users)
}

func selectNullRecord(db *gorm.DB) {
	//var user = entity.UserModel{Age: 0}
	//db.Debug().Where(user).Take(&user)
	//fmt.Println(user)

	// map可查空字段
	user := entity.UserModel{}
	db.Debug().Where(map[string]any{
		"age": 0,
	}).Take(&user)
	fmt.Println(user)
}

func manyWhereExample(db *gorm.DB) {
	var user entity.UserModel
	whereSQL := db.Where("age = ? and name = ?", 18, "张三")
	db.Debug().Where(whereSQL).Take(&user)
	fmt.Println(user)
}

func orExample(db *gorm.DB) {
	var users []entity.UserModel
	db.Debug().Or("age = 18").Or("name = ?", "小明").Find(&users)
	fmt.Println(users)
}

func notExample(db *gorm.DB) {
	var user entity.UserModel
	db.Debug().Not("age = 18").Take(&user)
	fmt.Println(user)
}

func sortExample(db *gorm.DB) {
	var userList []entity.UserModel
	// 降序
	db.Order("age desc").Find(&userList)
	fmt.Println(userList)
	// 升序
	db.Order("age asc").Find(&userList)
	fmt.Println(userList)
}

/*
查找特定字段
*/
func scanExample(db *gorm.DB) {
	var nameList []string
	db.Model(entity.UserModel{}).Select("name").Scan(&nameList)
	fmt.Println(nameList)
}

/*
查找特定字段
*/
func pluckExample(db *gorm.DB) {
	var nameList []string
	db.Model(entity.UserModel{}).Pluck("name", &nameList)
	fmt.Println(nameList)
}

/*
通过结构体返回数据
*/
func structExample(db *gorm.DB) {
	//type UserModel struct {
	//	Name string
	//	ID   int
	//}
	type UserModel struct {
		Type  string `gorm:"column:name"`
		Value int    `gorm:"column:id"`
	}
	var users []UserModel
	db.Debug().Model(UserModel{}).Scan(&users)
	fmt.Println(users)
}

func groupExample(db *gorm.DB) {
	type User struct {
		Age   int
		Count int
	}
	var users []User
	db.Model(entity.UserModel{}).Group("age").Select("age, count(*) as count").Scan(&users)
	fmt.Println(users)
}

/*
去重
*/
func distinctExample(db *gorm.DB) {
	var ageList []int
	db.Model(entity.UserModel{}).Distinct("age").Select("age").Scan(&ageList)
	fmt.Println(ageList)
}

func pageExample(db *gorm.DB) {
	var users []entity.UserModel
	//db.Limit(2).Offset(0).Find(&users)
	//fmt.Println(users)

	db.Limit(2).Offset(2).Find(&users)
	fmt.Println(users)
}

/*
Scope抽出可重复使用的方法
*/
func Age18(tx *gorm.DB) *gorm.DB {
	return tx.Where("age >= ?", 18)
}

func scopeExample1(db *gorm.DB) {
	var users []entity.UserModel
	db.Scopes(Age18).Find(&users)
	fmt.Println(users)
}

func nameIn(nameList ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("name in ?", nameList)
	}
}

func scopeExample2(db *gorm.DB) {
	var users []entity.UserModel
	db.Scopes(nameIn("张三", "李四")).Find(&users)
	fmt.Println(users)
}

/*原生sql*/
func rawExample(db *gorm.DB) {
	var users []entity.UserModel
	db.Raw("select * from users").Scan(&users)
	fmt.Println(users)
}

func execExample(db *gorm.DB) {
	db.Exec("update user_models set name = ? where id = ?", "张三丰", 1)
}

func main() {
	comm.Connect()
	execExample(comm.DB)
}
