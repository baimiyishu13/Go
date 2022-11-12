package main

import "fmt"

/*
+ 无参无返回值函数
+ 有一个参数的函数
+ 有两个参数的函数
+ 有一个返回值的函数
+ 有多个返回值的函数
*/
func main() {
	// 函数调用
	printinfo()
	myprint("有一个参数的函数")
	myprint(addstring("有两个参数的函数\t", "有一个返回值的函数"))
	x, y := swap("one", "two")
	fmt.Println(x, y)
}

// 无参无返回值函数
func printinfo() {
	fmt.Println("无参无返回值函数")
}

// 有一个参数的函数
func myprint(msg string) {
	fmt.Println(msg)
}

// 有两个参数的函数
func addstring(a, b string) string {
	c := a + b
	return c
}

// 有多个返回值的函数
func swap(x, y string) (string, string) {
	return y, x
}
