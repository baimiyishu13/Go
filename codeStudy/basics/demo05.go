package main

import "fmt"

func main() {
	var num int
	num = 1000
	fmt.Printf("num的值:%d,内存地址:%p\n", num, &num)

	num = 2000
	fmt.Printf("num的值:%d,内存地址:%p\n", num, &num)
}
