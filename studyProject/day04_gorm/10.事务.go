package main

import (
	"errors"
	"fmt"
	"studyProject/day04_gorm/comm"
	"studyProject/day04_gorm/entity"

	"gorm.io/gorm"
)

func automaticTransaction(db *gorm.DB) {
	var zhangsan, lisi entity.UserModel
	db.Take(&zhangsan, "name = ?", "张三")
	db.Take(&lisi, "name = ?", "李四")

	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&zhangsan).Update("money", gorm.Expr("money - 100")).Error
		err = errors.New("这是一个错误")
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = tx.Model(&lisi).Update("money", gorm.Expr("money + 100")).Error
		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func manualTransaction(db *gorm.DB) {
	var zhangsan, lisi entity.UserModel
	db.Take(&zhangsan, "name = ?", "张三")
	db.Take(&lisi, "name = ?", "李四")

	tx := db.Begin()
	err := tx.Model(&zhangsan).Update("money", gorm.Expr("money - 100")).Error
	err = errors.New("错误发生")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}

	err = tx.Model(&lisi).Update("money", gorm.Expr("money + 100")).Error
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}

	tx.Commit()
}

func main() {
	comm.Connect()
	manualTransaction(comm.DB)
}
