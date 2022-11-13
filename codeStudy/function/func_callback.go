package main

import "fmt"

func main() {
	r1 := add1(1, 2)
	fmt.Println(r1)

	r2 := oper(3, 4, add1)
	fmt.Println(r2)

	r3 := oper(10, 4, add1)
	fmt.Println(r3)

	// 既然可以调用实际函数，那么也可以调用匿名函数
	r4 := oper(12, 4, func(a int, b int) int {
		if b == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println(r4)
}

// 高阶函数,可以接受一个函数作为参数
func oper(a, b int, fun func(int, int) int) int {
	r := fun(a, b)
	return r
}

func add1(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
