package main

import "fmt"

func main() {
	// for 给控制变量赋值，循环条件，给控制变量增量或减量
	for i := 1; i <= 5; i++ {
		// 循环体
		fmt.Println("打印次数:", i)
	}
}
