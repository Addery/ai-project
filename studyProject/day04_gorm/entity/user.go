package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"not null;unique;size:20"`
	Age       int    `gorm:"not null;default:18"`
	Money     int    `gorm:"not null;default:0"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("BeforeCreate钩子函数执行了")
	//u.Name = "xiaoMing"
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("BeforeUpdate钩子函数执行了")
	return nil
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("BeforeDelete钩子函数执行了")
	return nil
}

func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind钩子函数执行了")
	return
}
