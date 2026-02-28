package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

// 一次性读取
func allReadOnly() {
	byteData, _ := os.ReadFile("day02/40文件操作/hello.txt")
	fmt.Println(string(byteData))
}

// GetCurrentFilePath 获取当前文件路径
func GetCurrentFilePath() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// 分片读取
func burstRead() {
	file, _ := os.Open("day02/40文件操作/hello.txt")
	defer file.Close()
	for {
		buf := make([]byte, 1)
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", buf)
	}
}

// 按行读
func laneRead() {
	file, _ := os.Open("day02/40文件操作/hello.txt")
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		fmt.Println(string(line))
		if err != nil {
			break
		}
	}
}

// 指定分割符
func splitRead() {
	file, _ := os.Open("day02/40文件操作/hello.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // 按照单词读
	//scanner.Split(bufio.ScanLines) // 按照行读
	//scanner.Split(bufio.ScanRunes) // 按照中文字符读
	//scanner.Split(bufio.ScanBytes) // 按照字节读读，中文会乱码

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//allReadOnly()
	//fmt.Println(GetCurrentFilePath())
	//burstRead()
	//laneRead()
	splitRead()
}
