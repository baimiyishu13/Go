package main

import "fmt"

// func() 本身就是一个数据类型；那么就可以定义一个函数类型的变量
func main() {
	// 不加括号，函数就是一个变量
	// f0()  如果加上了括号那就变成了函数
	fmt.Printf("%T\n", f0)  // func() | func(int, int) | func(int, int) int
	fmt.Printf("%T\n", 10)  // int
	fmt.Printf("%T\n", "A") // string

	// 定义函数类型的变量
	var f5 func(int, int)
	f5 = f0
	// 调用 f5
	f5(1, 2)
	fmt.Println(f5) // 0xbdfe40
	fmt.Println(f0) // 0xbdfe40

}
func f0(a, b int) {
	fmt.Println(a, b)
}
