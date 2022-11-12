package main

import "fmt"

func main() {
	str := "www.baidu.com"
	fmt.Println(str)

	// 获取字符串的长度 len
	fmt.Println("字符串的长度为:", len(str))

	// 获取指定的字节，默认从0开始
	//acsii 编码 119
	fmt.Println("字节打印:", str[0])
	fmt.Printf("%c", str[0])

	// for 遍历
	for i := 0; i < len(str); i++ {
		//fmt.Println(str[i])
		fmt.Printf("%c", str[i]) //wwww.baidu.com
	}

	// for range循环，便利数组、切片...
	// 返回下表和值
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c", v)
	}

	// string 不能修改的
	// cannot assign to str[2] (value of type byte)
	//str[2] = 'A'
	//fmt.Println(str[2])
}
