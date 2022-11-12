package main

import "fmt"

/* 打印一个方阵
* * * * * * 换行
* * * * * *
* * * * * *
* * * * * *
* * * * * *
 */
func main() {

	//for i := 0; i < 5; i++ {
	//	fmt.Print("* ")
	//}

	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
