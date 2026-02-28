package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	fmt.Println("Get Router")
	c.String(http.StatusOK, "Get Router")
}

func M1(c *gin.Context) {
	fmt.Println("M1 Request")
	c.Next()
	//c.Abort()
	fmt.Println("M1 Response")
}

func M2(c *gin.Context) {
	fmt.Println("M2 Request")
	//c.Next()
	c.Abort()
	fmt.Println("M2 Response")
}

func main() {
	r := gin.Default()

	r.GET("", M1, M2, get)

	fmt.Println("server running on :8080")
	r.Run(":8080")
}
