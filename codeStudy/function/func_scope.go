package main

import "fmt"

// 全局变量
var num int = 100

func main() {
	// 函数体内的局部变量
	temp := 100
	// if、for语句定义的一次性变量局部变量
	if b := 1; b <= 10 {
		// 语句内的局部变量
		temp := 50
		fmt.Println(temp) // 局部变量，就近原则
		fmt.Println(b)
	}
	fmt.Println(temp)
	fmt.Println(num)

	f1()
	f2()
}

// 内部
func f1() {
	a := 1
	fmt.Println(a)
	num := 30 // 就近原则
	fmt.Println(num)
}
func f2() {
	//fmt.Println(a)	// 不能使用其他函数定义的变量
	fmt.Println(num)
}
