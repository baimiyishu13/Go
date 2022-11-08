package main

import "fmt"

func main() {

	//分数
	var score int = 50

	// a b c d
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	}
	if score >= 80 && score < 90 {
		fmt.Println("B")
	}
	if score >= 70 && score < 80 {
		fmt.Println("C")
	}
	if score >= 60 && score < 70 {
		fmt.Println("D")
	}

	// 不需要写这么多if判断，使用else
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else if score >= 60 && score < 70 {
		fmt.Println("D")
	} else { //以上条件都不满足就走else
		fmt.Println("不及格")
	}
}
