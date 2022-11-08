package main

import "fmt"

func main() {
	var a int = 4
	var ptr *int

	ptr = &a
	fmt.Printf("ptr 的值为：%p\n", ptr) //ptr 的值为：0xc000020098
	fmt.Printf("ptr 为：%d\n", *ptr)  //ptr 为：4

}
