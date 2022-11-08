package main

import "fmt"

func main() {
	// var 定义变量，如果没有赋值，变量就默认是这个类型的默认值
	var (
		name string
		age  int
		addr string
	)
	name = "张三"
	age = 23
	addr = "shanghai"

	// strint默认值：空
	// int默认值：0
	fmt.Println(name, age, addr)
}
