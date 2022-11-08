package main

import "fmt"

var c int

func main() {
	// 声明局部变量 a和b 的和
	var c int = 3
	var b int = 4
	// 声明局部变量 c 并计算 a 和 b 的和
	c = c + b
	fmt.Printf("a=%d,b=%d,c=%d\n", c, b, c) //a=3,b=4,c=7
}
