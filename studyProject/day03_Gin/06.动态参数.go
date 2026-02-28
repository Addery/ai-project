package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println(id)
	})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
