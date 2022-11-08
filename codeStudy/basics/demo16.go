package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false

	// 逻辑 && 与；可以理解为 我和你都要满足
	// a 和 b都为真，true；一个为假结果为假；
	if a == true && b == true {
		fmt.Println(a, b)
	}
	fmt.Println(a && b) //false

	// 逻辑 ||或；一个真结果为真
	if a || b == true {
		fmt.Println(a, b) //true false
	}

	// 逻辑非 ！
	// 可以理解为取反，如果为假，结果为真
	fmt.Println(!a)
}
