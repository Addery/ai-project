package main

import (
	"fmt"
	"os"
)

func init() {
	// 读取配置文件中，结果路径错了
	_, err := os.ReadFile("xxx")
	if err != nil {
		panic(err.Error())
		//log.Fatalln("出错了")
	}
}

func main() {
	fmt.Println("啦啦啦")
}
