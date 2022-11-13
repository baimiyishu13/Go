package main

import "fmt"

// 匿名函数
func main() {
	f6()
	f7 := f6 // 函数本身也是一个变量
	f7()

	// 匿名函数
	f8 := func() {
		fmt.Println("f8函数")
	}
	f8()

	//简化；匿名哈桉树自己调用自己，只执行一次
	func() {
		fmt.Println("f9函数")
	}()

	//匿名函数传参
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//匿名函数传参 + return
	r1 := func(a, b int) int {
		fmt.Println(a, b)
		return a + b
	}(1, 2)
	fmt.Println(r1)
}

func f6() {
	fmt.Println("f6函数")
}
