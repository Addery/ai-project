package main

import (
	"fmt"
	"studyProject/day03_Gin/res"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"code": 0,
	//		"msg":  "ok",
	//	})
	//})

	r.GET("/ok", func(c *gin.Context) {
		res.Ok(c, res.DefaultCode, gin.H{}, "")
	})

	r.GET("/fail", func(c *gin.Context) {
		res.FailWithMsg(c, res.CodeMap[res.RoleErrCode])
	})

	fmt.Println("server is running at http://127.0.0.1:8080")
	r.Run(":8080")
}
