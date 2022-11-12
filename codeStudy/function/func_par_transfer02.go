package main

import "fmt"

// 引用传递
func main() {

	// 切片，可以扩容的数组
	s1 := []int{1, 2, 3, 4}
	fmt.Println("默认的数据:", s1) // [1 2 3 4]
	// 传入的是引用类型的数据，地址
	update2(s1)
	fmt.Println("调用后的数据:", s1) // [100 2 3 4]
}
func update2(s2 []int) {
	fmt.Println("传递的数据:", s2) // [1 2 3 4]
	s2[0] = 100
	fmt.Println("修改后的数据:", s2) // [100 2 3 4]
}
