package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type user1 struct {
	Name string `uri:"name"`
	Age  int    `uri:"age"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/:name/:age", func(c *gin.Context) {
		var u user1
		err := c.ShouldBindUri(&u)
		fmt.Println(u, err)
	})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
