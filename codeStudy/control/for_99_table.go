package main

import "fmt"

// 打印99乘法表
func main() {
	// 打印出 //0x9=0   1x9=9   2x9=18  3x9=27  4x9=36  5x9=45  6x9=54  7x9=63  8x9=72  9x9=81
	for i := 0; i < 10; i++ {
		fmt.Printf("%dx9=%d\t", i, i*9)
	}
	fmt.Println()

	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			fmt.Printf("%dx%d=%d\t", i, j, i*9)
		}
		fmt.Println()
	}

	// i < 10 替换成 i<j
	for j := 0; j < 10; j++ {
		for i := 0; i < j; i++ {
			fmt.Printf("%dx%d=%d\t", i, j, i*9)
		}
		fmt.Println()
	}
}
