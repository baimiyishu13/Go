package main

import "fmt"

func main() {
	// break 结束当前循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i)
	}
	fmt.Println()

	// continue 结束本次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
	}
}
