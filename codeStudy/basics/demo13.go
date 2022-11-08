package main

import "fmt"

func main() {
	var str string
	str = "Hello World"
	fmt.Printf("%T,%s\n", str, str) //string,Hello World

	//单引号和双引号的区别，单引号 字符，ascii字符码
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%d\n", v1, v1) //int32,65
	fmt.Printf("%T,%s\n", v2, v2) //string,A

	/*
		扩展知识
		中国的编码表
		Unicode编码表，号称兼容了全世界的文字
	*/
	v3 := '中'

	fmt.Printf("%T,%d\n", v3, v3) //int32,20013

	fmt.Println("zhang" + "san")

	fmt.Println("Hello\" Go")
	fmt.Println("Hello\n Go") // n 换行
	fmt.Println("Hello\t Go") // t 制表符
}
