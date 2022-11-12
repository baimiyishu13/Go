package main

import "fmt"

func main() {
	// 计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	i := 1
	for {
		fmt.Println(i)
		i++
		if i == 15 {
			break
		}
	}
}
