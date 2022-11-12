package main

import "fmt"

func main() {
	// 值传递
	// 传递，拷贝 arr
	// arr2 的数据是从arr1复制来的，所以是不同的空间
	// 修改 arr2 并不会影响 arr1
	// 值传递，传递的是数据的副本，修改数据，对于原始的数据没有影响
	// 值类型的数据，默认都是值传递，基础类型、array、strtuct
	// 定义一个数组 [个数]int
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	update(arr)
	fmt.Println("调用后的数据:", arr) // arr对象本身
	// 引用传递
}
func update(arr2 [4]int) {
	fmt.Println("arr2接受的数据:", arr2)
	arr2[0] = 100
	fmt.Println("arr2修改后的数据:", arr2)

}
