package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 多文件上传
	r.POST("/upMulFile", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, fileHeaders := range form.File {
			for _, fileHeader := range fileHeaders {
				c.SaveUploadedFile(fileHeader, "./upload/"+fileHeader.Filename)
			}
		}

		if err != nil {
			fmt.Println(err)
		}
	})

	r.POST("/upFile", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = c.SaveUploadedFile(fileHeader, "./upload/"+fileHeader.Filename)

		if err != nil {
			fmt.Println(err)
		}
	})

	//r.POST("/upFile", func(c *gin.Context) {
	//	fileHeader, err := c.FormFile("file")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	fmt.Println(fileHeader.Filename)
	//
	//	file, _ := fileHeader.Open()
	//	defer file.Close()
	//
	//	fileBytes, _ := io.ReadAll(file)
	//
	//	err = os.WriteFile(fileHeader.Filename, fileBytes, 0666)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//})

	fmt.Println("server run http://127.0.0.1:8080")
	r.Run(":8080")
}
