package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type user struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		var u user
		err := c.ShouldBindQuery(&u)
		fmt.Println(binding.Query.Name())
		fmt.Println(u, err)
	})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
