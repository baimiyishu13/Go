package main

import "fmt"

func main() {

	r1 := increment()
	fmt.Println(r1) // 地址：0x9de7e0

	v1 := r1()
	fmt.Println(v1) // 1
	v2 := r1()
	fmt.Println(v2)   // 2
	fmt.Println(r1()) // 3
	fmt.Println(r1()) // 4
	fmt.Println(r1()) // 5

	r2 := increment()
	fmt.Println(r2)   // 地址：0x9de7c0
	fmt.Println(r2()) // 1

}

// 自增
func increment() func() int {
	// 局部变量 i
	i := 0
	// 定义一个匿名函数，给变量自增并返回
	fun := func() int { // 内层函数，没有执行的
		i++
		return i
	}
	return fun
}
