package main

import "fmt"

/*
，
定义一个test()函数，它返回两个int类型的值，每次调用返回 100 和 200 俩个数值。
*/
func test() (int, int) {
	return 100, 200
}
func main() {
	// test() 函数返回两个值 100 和 200； 只需要第一个值，匿名函数废弃掉第二个值 即匿名函数 _
	a, _ := test()
	_, b := test()
	fmt.Println(a, b) //输出 100.200
}
