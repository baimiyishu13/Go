package main

import "fmt"

func main() {
	getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

// ... 可变参数
func getSum(nums ...int) {
	sum := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		sum += nums[i]
	}
	fmt.Println("sum:", sum)
}
