package main

import "fmt"

func main() {
	var a int = 10
	var b int = 3
	//var c int

	// + - * / % ++ --
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
}
