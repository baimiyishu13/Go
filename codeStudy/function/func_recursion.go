package main

import "fmt"

func main() {
	sum := getSum2(5)
	fmt.Println(sum)
}

// 5				5
// getSum2(4) + 5	15
// getSum2(3) + 4	10
// getSum2(2) + 3	6	getSum2(2) = getSum2(1) + 2
// getSum2(1) + 2	3	== 1 + 2
// 1				1
func getSum2(n int) int {
	if n == 1 {
		return 1
	}
	return getSum2(n-1) + n
}
