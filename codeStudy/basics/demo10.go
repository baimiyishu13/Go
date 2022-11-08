package main

import "fmt"

func main() {
	// var 变量名 数据类型
	// bool: true false(默认值)
	var isFlage1 bool        // false
	var isFlage2 bool = true //trye

	fmt.Printf("%T,%t\n", isFlage1, isFlage1) //bool,false
	fmt.Printf("%T,%t\n", isFlage2, isFlage2) //bool,true
}
