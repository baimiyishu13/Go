package main

import (
	"fmt"
)

var name string = "lisi"

func main() {
	//局部变量
	var name string = "zhangsan"
	var age1 int = 18
	fmt.Println(name, age1)
}

func test1() {
	fmt.Println(name)
}
