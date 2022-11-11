package main

import "fmt"

// if 嵌套语句验证密码
func main() {

	// 验证密码，再次输入密码
	var a, b string
	var pwd string = "go@123"
	// 用户输入
	fmt.Print("请输入密码:")
	fmt.Scanln(&a)
	if a == pwd {
		fmt.Print("再次输入密码:")
		fmt.Scanln(&b)
		if b == pwd {
			fmt.Println("登录成功")
		}
	} else {
		fmt.Println("密码错误，请重新登录")
	}
}
