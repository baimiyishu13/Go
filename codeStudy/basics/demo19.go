package main

import "fmt"

func main() {
	var a int = 21
	var b int

	// =赋值
	b = a
	fmt.Println(b) //21

	// += 加等 b = b + a; 简化代码
	b += a
	fmt.Println(b) //42
}
