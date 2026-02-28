package comm

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm_new_db?charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect err: %v", err)
	}
	DB = db
}
