package main

import "fmt"

func main() {
	a := 5.0    // float
	b := int(a) // 转为int

	//需求：将int类型的 a 转换成位 float64 类型 类型转换
	c := float64(a)
	fmt.Printf("%T\n", c) //float64

	d := byte(a)
	fmt.Printf("%T,%d\n", d, d) //uint8,5

	fmt.Printf("%T,%f\n", a, a) //float64,5.000000
	fmt.Printf("%T,%d\n", b, b) //int,5
}
