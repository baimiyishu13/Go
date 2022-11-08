package main

import (
	"bufio"
	"fmt"
	"os"
)

// 多行声明
type User struct {
	Username string
	Email    string
}

func main() {
	u := User{
		Username: "baby",
		Email:    "baby@123.com",
	}
	fmt.Println(u)
	/*
		输入和输出：fmt包

		输出：
			fmt.Println()	//打印换行
			fmt.Print()  	 //打印
			fmt.Printf()	//格式化打印
		输入：
			fmt.Scanf()
			fmt.Scanln()
			fmt.Scan
	*/

	var x int
	var y float64

	fmt.Println("请输入一个整数，一个浮点数：")
	//读取键盘的输入，通过操作地址，赋值给x和y
	//变量取地址  &变量
	//Scanln 阻塞等待你的键盘输入
	fmt.Scanln(&x, &y)
	fmt.Printf("x:%d,y:%f\n", x, y)

	//在未来的IO中，可以使用 bufio 包来读取输入的内容
	fmt.Println("请输入一串字符串:")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读出的内容为:", s1)

}
