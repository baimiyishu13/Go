package main

import "fmt"

func main() {
	a := false
	switch a {
	case false:
		fmt.Println("1.case的条件为false")
		fallthrough // case穿透，不管下一条条件是否满足，都会执行
	case true:
		if a == false {
			break //终止穿透
		}
		fmt.Println("2.case的条件为true")
	}
}
