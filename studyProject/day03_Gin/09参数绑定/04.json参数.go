package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type user3 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var u user3
		err := c.ShouldBindJSON(&u)
		fmt.Println(u, err)
	})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
