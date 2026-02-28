package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type user4 struct {
	Name        string `header:"Name"`
	Age         int    `header:"Age"`
	UserAgent   string `header:"User-Agent"`
	ContentType string `header:"Content-Type"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var u user4
		err := c.ShouldBindHeader(&u)
		fmt.Println(u, err)
	})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
