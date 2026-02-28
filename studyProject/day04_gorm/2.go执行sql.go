package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func exec() {
	db, connErr := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm_new_db")
	if connErr != nil {
		log.Fatalf("数据库连接失败%v", connErr)
	}

	defer db.Close()

	// 除了查询以外，使用Exec函数
	_, dbErr := db.Exec("CREATE TABLE users(id INT NOT NULL , name VARCHAR(20), PRIMARY KEY(id));")
	if dbErr != nil {
		log.Fatalf("执行sql语句时出错%v", dbErr)
	}

	fmt.Println("建表成功")
}

func query() {
	db, connErr := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm_new_db")
	if connErr != nil {
		log.Fatalf("数据库连接失败%v", connErr)
	}

	defer db.Close()

	//rows, err := db.Query("select * from users")
	//if err != nil {
	//	log.Fatalf("查询时出错%v", err)
	//}
	//for rows.Next() {
	//	var id int
	//	var name string
	//	err = rows.Scan(&id, &name)
	//	fmt.Printf("id:%d name:%s err:%v\n", id, name, err)
	//}

	/*
		查一行
	*/
	var id int
	var name string
	err := db.QueryRow("select * from users").Scan(&id, &name)
	fmt.Printf("id:%d name:%s err:%v \n", id, name, err)

}

func main() {
	//exec()
	query()
}
