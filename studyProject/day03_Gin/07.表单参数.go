package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		age, b := c.GetPostForm("age")
		fmt.Printf("%#v\n", name)
		fmt.Printf("%#v%v\n", age, b)
	})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
