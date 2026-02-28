package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm_new_db")
	if err != nil {
		log.Fatalf("数据库连接失败%v", err)
	}
}
