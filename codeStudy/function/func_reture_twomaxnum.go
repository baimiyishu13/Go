package main

func main() {

	// 形参与实参要一一对应，顺序，个数，类型
	max(1, 2)
}

// max 两个数字比大小
// 形式参数：函数定义时，用来接受外部传入数据的参数，就是形式参数
// 实际参数：程序实际调用时，传入的参数
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	// 一个函数上定义了返回值，那么函数中必须使用 return 语句
	// 返回值
	// 调用处需要使用变量接受该结果
	return result
}
