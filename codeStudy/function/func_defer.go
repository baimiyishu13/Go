package main

import "fmt"

func main() {
	f("1")
	fmt.Println("2")
	defer f("3") // 会被延迟到最后执行
	fmt.Println("4")
	defer f("5")
	defer f("6") //defer语句会按照逆序执行
}

func f(s string) {
	fmt.Println(s)
}
