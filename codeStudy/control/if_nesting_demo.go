package main

import "fmt"

// if 语句嵌套
func main() {

	var a int = 100
	var b int = 200

	if a == 100 {
		fmt.Println("a 满足条件")
		if b == 200 {
			fmt.Println("b 满足条件")
		}
	}
}
