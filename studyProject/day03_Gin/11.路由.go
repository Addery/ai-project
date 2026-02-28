package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMethod(c *gin.Context) {
	url := c.Request.URL
	fmt.Println(url)
}

func UserGroup(r *gin.RouterGroup) {
	r.GET("users", TestMethod)
	r.POST("users", TestMethod)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	group := r.Group("api")
	UserGroup(group)

	fmt.Println("server running on :8080")
	r.Run("localhost:8080")
}
