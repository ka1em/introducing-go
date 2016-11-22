package main

import (
	"fmt"
	"os"
)

func main() {

	//打开文件
	file, err := os.Open("../ch01/env.sh")
	if err != nil {
		fmt.Println("err")
		return
	}

	//关闭
	defer file.Close()

	//获取文件属性，文件大小
	stat, err := file.Stat()
	if err != nil {
		return
	}
	bs := make([]byte, stat.Size())

	//读文件
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
