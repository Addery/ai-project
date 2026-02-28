package main

import (
	"fmt"
	"io"
	"os"
)

func writerFileOnly() {
	err := os.WriteFile("day02/40文件操作/file1.txt", []byte("这是内容"), os.ModePerm)
	fmt.Println(err)
}

func copyFile() {
	rFile, _ := os.OpenFile("day02/40文件操作/file1.txt", os.O_RDONLY, 0777)
	wFile, _ := os.OpenFile("day02/40文件操作/file2.txt", os.O_CREATE|os.O_WRONLY, 0777)
	defer rFile.Close()
	defer wFile.Close()
	n, err := io.Copy(wFile, rFile)
	fmt.Println(n, err)
}

func scanDir() {
	dir, _ := os.ReadDir("day02/40文件操作")
	for _, d := range dir {
		info, _ := d.Info()
		fmt.Println(d.Name(), info.Size(), info.ModTime())
	}
}

func main() {
	//writerFileOnly()
	//copyFile()
	scanDir()
}
