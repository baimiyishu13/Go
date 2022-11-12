package main

import "fmt"

// main 函数
func main() {
	fmt.Println("hello,world")

	fmt.Println(add(1, 2))
}

/*
	func 函数名(参数,参数...) 函数调用后的返回值 {
		函数体   //执行一段代码
		reture	//返回结果
	}
*/
func add(a, b int) int {
	c := a + b
	return c
}
