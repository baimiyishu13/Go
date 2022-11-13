package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a=", a)
	defer f3(a) // 函数中的a= 10;参数以及传递进入，在最后执行
	a++
	fmt.Println("a=", a)

}
func f3(s int) {
	fmt.Println("函数中的a=", s)
}
