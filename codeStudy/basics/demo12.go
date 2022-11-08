package main

import "fmt"

func main() {
	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 5.12
	// %f 默认保留小6位数，%.2f就是暴力两位，%.3f就是暴力3位
	fmt.Printf("%T,%f\n", f1, f1)
	fmt.Printf("%T,%.3f\n", f2, f2)

	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	//num1 = -123.00009 num2 = -123.0000901
	fmt.Println("num1 =", num1, "num2 =", num2)
}
