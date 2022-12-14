## :bell:学前准备工作

#### :trophy: Hello，World

记事本编写 hello.go


```go
package main

import "fmt" //导入一个系统包"fmt"用来输出的

func main() {
	fmt.Println("Hello,World!") //答应输出Hello,World!字符串
}
```

当前目录下执行：`go run hello.go`

 ![image-20221105165951801](/images/image-20221105165951801.png)

**解释：**

在GO语言中命名为`mian`的包具有特殊的含义，GO语言的编译程序会试图把这种名字的包编译成二进制的可执行的包。所以使用GO语言编译的可执行程序都必须包含一个名交`main`的包，一个可执行的程序有且只有一个`main`包。

当解释器发现某个包的包名为`main`时，它一定也会发现名未`mian()`的函数，否则不会创建可执行文件。

`main()`函数是程序的入口，如果没有这个程序，程序就没办法开始执行。程序编译时，会使用声明`main`包的代码所在的目录的目录名作为二进制可执行文件的文件名。

设置go111module（环境变量）=off

> 准备工作完成！



:bell:



:trophy:











## :bell:基础语法学习



### :trophy: 注释

我们在未来写代码的时候会有一种情况时常发生，刚写完的代码，觉得逻辑清晰，自己怎么这么厉害，然后过了一个月再去看这个代码，就会产生疑问，这个代码到底TM谁写的！所以为了防止自己看不懂自己写的代码。或者你想把你写的代码给别人看，这个时候，我们就需要注释了！

> 注释：你可以理解为写给别人看的，作为一个开源项目。机器不会去执行这些语句。

**注释主要的功能就是为了增强代码的可读性，不参与程序的一切功能，Go语言的注释主要分为两类。**

1. 单行注释
2. 多行注释

```GO
package main

import "fmt"

// 我是单行注釋，程序不会执行，写给自己看的

/*
我是多行注释，使用杠星，星表示在这个区间可以换行写多行注释，也十分常用！
这是一个main函数，这个是go语言的入口
*/

func main() {
	// fmt.Println 打印一句话，然后执行完毕后进行换行
	fmt.Println("Hello,World!")
}
```

写好注释是一个非常良好的习惯，我们都应该按照要求给自己代码写好注释，为了自己，为了他人。很多大公司也是对注释有严格的要求。



### :trophy:变量

搞清楚注释后，接下来看一个程序中十分重要的东西，变量！

在数学的概念中，变量表示没固定值且可以改变的。比如x=1，y=2；

字面上的意思理解：变量就是会变化的量，比如我定了一个变量叫做名字；它在Go语言中是这样表示的：这个值可以是张三，也可以是李四，也可以是王五。那么在这里，这个`name`就是变量。可以变化的量。

```GO
//这里的 “=” 是赋值，就是吧等号右边的值，赋值给左边的意思。
package main

import "fmt"

func main() {

	// name 变量
	var name string = "tom"
	fmt.Println(name)

	name = "zhangsan"
	fmt.Println(name)
}
```

**解释：**

​	内存空间-地址编号[0x11112345] 存了一个字符串 `tom`；为这个内存空间定义了一个名字叫`name`；通过这个名字指向了内存空间；空间中的`tom`可以理解未一个盒子，可以拿掉换成`zhangsan`的盒子；

`name`：代表的是一个变量，变量对应的内存底下是串内存地址空间，数据放在空间中；

再次印证了：变量就是会变化的量



#### :four_leaf_clover:变量的定义

Go语言是静态类型语言，就是所有的类型我们都需要明确的定义，我们这里先不管其他类型，先了解string，我们用它来表示字符串！

在Go语言中，声明了一个变量一般是使用var关键字：

```GO
var name type
```

+ 第一个var：声明变量的关键字，是固定写法，要声明一个变量就需要一个var。
+ 第二个name：就是我们变量的名字，按照需求起一个名字，用来后续使用！
+ 第三个type，就是用来表示变量的类型。

举个例子：

```GO
// 定义一个字符串变量 name
var name string

// 定一个数字类型变量 age
var age int

```

Go和其他许多编程语言不通同，他在声明变量时间变量的类型放在变量名之后，这样做的好处就是可以避免像C语言中那样含糊不清的声明形式，例如：`int* a,b;` 其中只有是指针而b不是。如果想要这两变量都是指针，则需要将它们分开书写。而在GO中，则可以很简单的将其声明指针类型：

```go
var a,b *int
```

变量的明明规则遵循骆驼命名法，即每个新单词首个字母大写，例如：userFile和 systemInfo。

> 定义变量的标准格式为

```
var 变量名 变量类型
```

变量声明以关键字开头，后置变量类型，行尾无需分号！

> 我们有时候会批量定义变量，如果每次单独定义会很麻烦，Go语言支持批量定义变量

是有关键字var和括号，可以将变量定义放在一起。

```GO
func main() {
    // var 定义变量，如果没有赋值，变量就默认是这个类型的默认值
	var (
		name string
		age  int
		addr string
	)
	// strint默认值：空
	// int默认值：0
	fmt.Println(name, age, addr)
}
```

var 形式的声明语句 往往是用于需要显式指定变量类型的地方，或者因为变量稍后会被重新赋值而初始值无关紧要地方。

当一个变量被声明后，如果没有显示的给它赋值，系统会自动赋予它该类型的零值

+ 整型和浮点型的默认值是`0`  和`0.0`
+ 字符串变量的默认值为`空字符串`
+ 布尔型变量默认值为`false`
+ 切片、函数、指针变量默认值为`nil`



#### :four_leaf_clover:变量的初始化

> 变量初始化的标准格式

```go
var 变量名 类型 = 值(表达式)
```

比如，想定义`zhangsan`的一些信息，可以表示

```GO
func main() {
	var name1 string = "zhangsan"
	var age1 int = 18
	fmt.Printf("name:%s,age:%d", name1, age1)
}
```

这里的name和age就是变量名，name的类型为string，age的类型为int，他们的值分别为`张三`和`18`

> 短变量声明初始化
>

```go
// := 自动推到
name2 := "zhangsan"
age2 := 18
```

这是Go语言的推导声明写法，编译器会自动根据右值推断出左值的类型。

它可以自动推导出一些类型，但是使用也是有限制的；

+ 定义变量，同时显式初始化
+ 不能提供数据类型
+ 只是用在函数内部，不能随便到处定义

因为简洁和灵活的特点，简短变量声明被广泛用于大部分的局部变量的声明和初始化。

注意：由于使用了 `:=`, 而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。

```GO
// 定义变量 name
var name string
// 定义变量 name，并赋值“zhangsan”
name := "zhangsan"
```

编译会报错，信息如下：

```GO
.\demo04.go:14:7: no new variables on left side of :=
```



#### :four_leaf_clover:理解变量（内存地址）

Printf：

+ %d：打印数字类型
+ %p：打印内存地址 (取地址符 &变量名)

```GO
	var num int
	num = 1000
	fmt.Printf("num的值:%d,内存地址:%p\n", num, &num)
    //num的值:1000,内存地址:0xc000020098

	num = 2000
	fmt.Printf("num的值:%d,内存地址:%p\n", num, &num) //取地址符 &变量名
	//num的值:2000,内存地址:0xc000020098
```



#### :four_leaf_clover:变量的交换

传统语言交换方式

```c++
// 在编程中，最简单的就是内存交换了，但是我们一般常见的方式就是定义中间变量
a = 100
b = 200
temp = 0
temp = a  // 第一步a存成了temp
a = b     // 第二步b赋值给a
b = temp  // 第三步temp赋值给b
```

Go语言有个十分简单的写法：

```GO
func main() {
    var a int = 100
	var b int = 200
	 
    // 在Go 语言中，可以直接这样来实现换值，不需要中间变量了。
	b, a = a, b
	fmt.Println(a, b)
}
```



#### :four_leaf_clover:匿名变量

匿名变量的特点是一个下划线`_`,`_`本身就是一个特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋值给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。使用匿名变量时，只需要在变量声明的地方是有下划线代替即可，例如

```GO
package main

import "fmt"
/*，
定义一个test()函数，它返回两个int类型的值，每次调用返回 100 和 200 俩个数值。
 */
func test() (int, int) {
	return 100, 200
}
func main() {
	// test() 函数返回两个值 100 和 200； 只需要第一个值，匿名函数废弃掉第二个值 即匿名函数 _
	a, _ := test()
	_,b := test()
	fmt.Println(a,b) //输出 100.200
}
```

编码过程中，可能遇到没有名称的变量名、类型、方法。虽然这不是必须的，但是有时候这样做可以极大地增强代码的灵活性，这些变量被统称为匿名变量。

匿名变量不占用内存空间，不会分配内存，匿名变量与匿名变量之间也不会因为多次声明而无法使用



#### :four_leaf_clover:匿名的作用域

**一个变量（常量、类型或函数）在程序中都有一定的作用范围，称之为：作用域**

了解变量的作用域对学习Go语言是比较重要的，因为Go语言会在编译时检查么个变量是否使用过，一旦出现未使用的变量，就会报编译错误。如果不能理解变量的作用域，就有可能会带来一些不明所以的编译错误。

> 局部变量

在函数体内声明的变量称之为局部变量，它们的作用域只有在函数体内，函数的参数和返回值变量都属于局部变量。

```go
func main() {
	// 声明局部变量 a和b 的和
	var a int = 3
	var b int = 4
	// 声明局部变量 c 并计算 a 和 b 的和
	c := a + b
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c) //a=3,b=4,c=7
}
```

> 全局变量

在函数体外声明的变量，全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，当然，不包含这个局部变量的源文件需要使用`import` 关键字引入全局变量所在的文件源文件之后才能使用这个全局变量。

全局变量声明必须以var关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。

```GO
// 声明全局便量
var c int
func main() {
	// 声明局部变量 a和b 的和
	var a int = 3
	var b int = 4
	// 声明局部变量 c 并计算 a 和 b 的和
	c = a + b
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c) //a=3,b=4,c=7
}
```

Go语言程序中全局变量和局部变量 名称可以相同，但是函数体内的`局部变量`会被优先考虑。

```gO
// 声明全局变量
var a float64 = 3.14

func main() {
	//声明局部变量
	var a int = 3
	fmt.Println(a) // a = 3
}
```



#### :four_leaf_clover:常量

**常量是一个简单值的标识符，在程序运行时，不会被修改的量**

注意：常量中的数据值可以是布尔值、数字型（整数型、浮点型和复数）和字符串型。

```go
const identifier [type] = value
```

可以省略类型说明符，因为编译器可以根据变量的值来推断其类型。

+ 显式类型定义：`const b string = “abc”`
+ 隐式类型定义：`Const b = "abc"`

多个相同类型的声明可以简写为：

```SH
const c_name1, c_name2 = value1, value2
```

常量的定义：const

```GO
func main() {
	const url string = "www.baidu.com"     //显式定义
	const url2 = "www.baidu.com"           //隐式定义
	const a, b, c = 3.14, "zhangsan", true //同时定义多个常量

	fmt.Println(url)
	fmt.Println(url2)
	fmt.Println(a, b, c) //3.14 zhangsan true
}
```



#### :four_leaf_clover:常量iota

中文名是“埃欧塔”，计数器

iota：特殊常量，可以认为是一个可以被编译器修改的常量，iota是go语言的**常量计数器**

iota在const关键字出现时将被重置为0 （const 内部的第一行之前），const中每新增一行常量声明将使iota可理解为const语句块中的索引。

iota可以被用作枚举值：

```GO
	const (
		a = iota
		b = iota
		c = iota
	)
```

第一个iota等于0，每当iota在新的一行被使用时，它的值都会自动加1；所以a=0,b=1,c=2 可以简写如下形式：

iota无论走不走都会进行计数，直到它的恢复！

```GO
	const (
		a = iota	// 0
		b			// 1
		c			// 2
		d = "haha"	// haha iota 3
        e			// haha iota 4 (e 如果不赋值，默认 haha)
        f = iota  	// 5 	iota 5	(恢复iota计数)
		j			// 6	iota 6
	)
	//新的一组会恢复计数
	const (
		n = iota	// 0
		m			// 1
	)
```

例子：



### :trophy: 基本数据类型

Go语言是一种静态类型的编程语言，在Go编程语言中，数据类型用于声明声明函数和变量。数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存就可以充分利用内存。编辑器在进行编译的时候，就要知道每个值的类型，这样编译器就知道要为这个值分配多少内存，并且知道这段分配的内存表示什么。

 ![image-20221106201935065](/images/image-20221106201935065.png)

#### :four_leaf_clover:布尔型

布尔型的值只有两个 `true`和`false`，一个简单的例子：var b bool = true、

```go
func main() {
   // var 变量名 数据类型
   // bool: true false(默认值)
   var isFlage1 bool        // false
   var isFlage2 bool = true //trye

   fmt.Printf("%T,%t\n", isFlage1, isFlage1) //bool,false
   fmt.Printf("%T,%t\n", isFlage2, isFlage2) //bool,true
}
```

#### :four_leaf_clover:数值型

整形 int 和 浮点型 float32、float64 ，Go语言支持整形和浮点型数字，并且支持复数，其中位的运算采用补码。

> Go也有基于框架结构的类型，例如：uint无符号，int有符号

| 序号 | 类型 & 描述                                       |
| ---- | ------------------------------------------------- |
| 1    | uint8 无符号 8 位整数（0到255）                   |
| 2    | uint16 无符号 16 位整数（0到65535）               |
| 3    | uint32 无符号 32 位整数（0到4229496795）          |
| 4    | uint64 无符号 64位整数（0到18446744073709551615） |
| 5    | int8 有符号 8 位整数（-128-127）                  |
| 6    | int16 有符号 16 位整数（-128-127）                |
| 7    | int32 有符号 23 位整数（-128-127）                |
| 8    | int64 有符号 46 位整数（-128-127）                |

```go
func main() {
   // 定义一个整形
   var age int = 8
   fmt.Printf("%T,%d\n", age, age) //int,8
}
```

> 浮点型

| 浮点型 | 类型 & 描述                   |
| ------ | ----------------------------- |
| 1      | float32 IEEE-754 32位浮点型数 |
| 2      | float64 I-EEE754 64位浮点型数 |
| 3      | complex64 32位实数和虚数      |
| 4      | complex64 64位实数和虚数      |

```go
func main() {
   var f1 float32
   f1 = 3.14
   var f2 float64
   f2 = 5.12
   // %f 默认保留小6位数，%.2f就是暴力两位，%.3f就是暴力3位
   fmt.Printf("%T,%f\n", f1, f1)
   fmt.Printf("%T,%.3f\n", f2, f2)
}
```

1. 关于浮点数在机器中存放形式的简单说明：**浮点数=符号位+指数位+尾数位**
2. 尾数数部分可能丢失，造成精度损失。-123.0000901

```go
var num1 float32 = -123.0000901
var num2 float64 = -123.0000901
//num1 = -123.00009 num2 = -123.0000901
fmt.Println("num1 =", num1, "num2 =", num2)
```

说明：

1. float64 精度要比float32 准确
2. 保存高精度数值，应该使用float64

浮点型的存储分为三个部分：**符号位+指数位+尾数位**，在存储过程中，精度可能会有丢失

goland的浮点型默认`float64类型`

通常情况下，应该使用float64，因为它比float32更精确

> 以下列出了其他更多的数字类型：

| 序号 | 类型 & 描述                          |
| ---- | ------------------------------------ |
| 1    | byte 类似 uint8                      |
| 2    | rune 类似 int32                      |
| 3    | uint32 或64位                        |
| 4    | int 与 uint 一样大小                 |
| 5    | uintptr 无符号整型，用于存放一个指针 |



#### :four_leaf_clover:字符串型


字符串就是一串固定长度的字符连接起来的字符序列。

Go的字符串是由单个字符连接起来的，Go语言的字符串的字节使用UTF-8编码标识 Unicode文本。

```go
package main

import "fmt"

func main() {
	var str string
	str = "Hello World"
	fmt.Printf("%T,%s\n", str, str) //string,Hello World

	//单引号和双引号的区别，单引号 字符，ascii字符码
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%d\n", v1, v1) //int32,65
	fmt.Printf("%T,%s\n", v2, v2) //string,A

	/*
		扩展知识
		中国的编码表：GBK
		全世界的编码表：Unicode编码表，号称兼容了全世界的文字
	*/
	v3 := '中'
	fmt.Printf("%T,%d\n", v3, v3) //int32,20013
}
```

Go语言的字符串连接可以通过 + 实现：

```Go
fmt.Println("zhang" + "san") //zhangsan 
```

转义字符：

```go
fmt.Println("Hello\" Go")
fmt.Println("Hello\n Go") // n 换行
fmt.Println("Hello\t Go") // t 制表符
```



#### :four_leaf_clover:数据类型转换

再必要以及可执行的情况下，一个类型的值可以被转换成另一种类型的值，由于Go语言不存在隐式类型转换。因此所有的类型转换都必须显式声明：

```go
valueofTypeB = typeB(valueofTypeA)
```

类型B的值 = 类型B(类型A的值)

```go
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
```

类型转换只能再定义正确的情况下转换成功，例如：

+ 冲一个取值范围小的类型，转换成一个取值范围大的类型（int16 转成 int32）

当从一个取值范围较大的类型转到取值范围较小的类型时，（int32 转 int16 或者其他），会发生精度丢失（截断）的情况！



### :trophy:运算符

+ 算数运算符
+ 关系运算符
+ 逻辑运算符
+ 位运算符
+ 赋值运算符
+ 其他运算符

 ![image-20221108164249728](/images/operator.png)

#### :four_leaf_clover:算数运算符

下表列出了所有Go语言的算数运算符，假定A值位10，B值位20

| 运算符 | 描述 | 实例                  |
| ------ | ---- | --------------------- |
| +      | 相加 | A + B 输出结果为30    |
| -      | 相减 | A - B 输出结果位为-10 |
| *      | 相乘 | A * B 输出结果为 200  |
| \      | 相除 | B \ A 输出结果为 2    |
| %      | 求余 | B % A输出结果为 0     |
| ++     | 自增 | A++ 输出结果为 11     |
| --     | 自减 | A-- 输出结果为9       |

```go
func main() {
   var a int = 10
   var b int = 3
   //var c int

   // + - * / % ++ --
   fmt.Println(a + b)
   fmt.Println(a - b)
   fmt.Println(a * b)
   fmt.Println(a / b)
   a++	// a = a+1
   fmt.Println(a)
   a--	// a = a-1
   fmt.Println(a)
}
```



#### :four_leaf_clover:关系运算符

下表列出了所有Go语言的关系运算符，假定A值为10，B值为20.

返回：True、False

| 运算符 | 描述                                                |                   |
| ------ | --------------------------------------------------- | ----------------- |
| ==     | 检查两个值是否相等，如果是返回True负责返回False     | (A == B) 为 False |
| !=     | 检查两个值是否不相等，如果是返回True负责返回False   | A != B) 为 True   |
| >      | 检查左边值是否大于右边，如果是返回True负责返回False | A > B) 为 False   |
| <      | 检查左边值是否小于右边，如果是返回True负责返回False | A < B) 为 True    |
| >=     | 检查左边值是否大于右边，如果是返回True负责返回False | A >= B) 为 False  |
| <=     | 检查左边值是否小于右边，如果是返回True负责返回False | A <= B) 为 False  |

```go
func main() {
   var a int = 11
   var b int = 10
   // 关系运算符返回的都是布尔值

   // 判断 if 如果 ... 结果
   if a > b {
      // 执行一些 a>b 操作
   } else { //else 否则
   
   }
}
```



#### :four_leaf_clover:逻辑运算符

下表列出了所有Go语言的逻辑运算符。加入A值为True，B值为False

> 2件事判断 18岁 和 身份证，满足两个才能进入网吧

| 运算符 | 描述                                                         | 实例             |
| ------ | ------------------------------------------------------------ | ---------------- |
| &&     | 逻辑 and 运算符，如果两边的操作都是True，则条件True，否则为False | (A && B)为False  |
| \|\|   | 逻辑 or 运算符，如果两边的操作有一个True，则条件True，否则为False | (A \|\| B)为True |
| ！     | 逻辑 not 运算符，如果条件为TrueT，则逻辑not条件false，否则为True | !(A && B)为True  |

```go
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
```



#### :four_leaf_clover:位运算符

> 二进制

Go语言支持的位运算符如下，假设A为60，B为13

60二进制：11 1100

13二进制： 00 1101

| 运算符 | 描述                                                         | 实例                            |
| ------ | ------------------------------------------------------------ | ------------------------------- |
| &      | 按位与运算符“&”是双目运算符，都是1结果位1，否则是0           | (A & B) 结果12，001100          |
| \|     | 按位或运算符“\|”是双目运算符，都是0结果位0，否则是1          | (A \| B) 结果61,111101          |
| ^      | 按位异或运算符“^”是双目运算符，不同位1，相同为0              | (A ^ B)结果49，110001           |
| &^     | 位清空，a&^b,对于b上的每个数值，如果为0，则取a对应位上的数值，如果为1，则取0 | (A &^ B)结果48，110000          |
| <<     | 左移运算符“<<”是双目运算，左移n位就是乘以2的n次方，其功能是吧"<<"左边的运算数的各二进制位全部左移若干位，由“<<”右边的数指定移动到的位数。 | A<<2结果240（60*2^2），11110000 |
| >>     | 右移运算符">>"是双目运算，右移n位就是除以2的n次方，其功能是吧">>"右边的运算数的各二进制位全部右移若干位，由“>>”左边的数指定移动到的位数。 | A>>2结果15（60/4），1111        |

用途：正常情况下很少使用，再底层加解密会使用；将所有的位向左/右移几位，打开后就是乱码了；底层都是二进制；另外一个程序在执行前恢复就完成了解密加密。

底层：高低电频（电路问题）- 二进制汇编

%b：打印二进制

```go
func main() {
   /*
      二进制 0 和 1，逢二进一
      位运算：二进制上的  0：false 1：true
      & : 都是1结果为1，否则是0
      | : 都是0结果为0，否则是1
      60: 0011 1100
      13: 0000 1101
      -------------
      &  0000 1100  同时满足
      |  0011 1101  一个满足
      ^   0011 0001  不同为1，相同为0（异或）
   */
   var a uint = 60
   var b uint = 13
   //位运算
   var c uint
   c = a & b                       //位运算
   fmt.Printf("%d,二进制：%b\n", c, c) //12,二进制：1100;%.8d显示8位

   c = a | b                       //位运算
   fmt.Printf("%d,二进制：%b\n", c, c) //61,二进制：111101

   c = a ^ b                       //位运算
   fmt.Printf("%d,二进制：%b\n", c, c) //49,二进制：110001

   c = a >> 2                      //位运算
   fmt.Printf("%d,二进制：%b\n", c, c) //15,二进制：1111

   c = a << 2                      //位运算
   fmt.Printf("%d,二进制：%b\n", c, c) //240,二进制：11110000
}
```





#### :four_leaf_clover:赋值运算符

下表列出了所有G语言的赋值运算符

| 运算符 | 描述                                             | 实例                 |
| ------ | ------------------------------------------------ | -------------------- |
| =      | 简单的赋值运算符，将一个表达式的值赋值给一个左值 | c = a + b            |
| +=     | 相加后再赋值                                     | c +=a 等于 c =c +a   |
| -=     | 相减后再赋值                                     | c -=a 等于 c =c -a   |
| *=     | 相乘后再赋值                                     | c *=a 等于 c =c *a   |
| /=     | 相除后再赋值                                     | c /=a 等于 c =c /a   |
| %=     | 求余后再赋值                                     | c %=a 等于 c =c %a   |
| <<=    | 左移后赋值                                       | c <<=2 等于 c = c<<2 |
| >>=    | 右移后赋值                                       | c >>=2 等于 c = c>>2 |
| &=     | 按位与后赋值                                     | c&=a 等于 c = c&a    |
| ^=     | 按位异或后赋值                                   | c^=a 等于 c = c^a    |
| \|=    | 按位或后赋值                                     | c\|=a 等于 c = c\|a  |

重点理解上半部分，下半部分不常用

```go
func main() {
   var a int = 21
   var b int

   // =赋值
   b = a
   fmt.Println(b) //21

   // += 加等 b = b + a; 简化代码
   b += a
   fmt.Println(b) //42
}
```



#### :four_leaf_clover:其他运算符

下表列出来Go语言的其他运算符

| 运算符 | 描述               | 实例                     |
| ------ | ------------------ | ------------------------ |
| &      | 返回变量存储的地址 | &a；将给出变量的实际地址 |
| *      | 指针变量           | *a；是一个指针变量       |



### :trophy:拓展:键盘输入输出

程序需要交互

 ![image-20221108175556816](/images/image-20221108175556816.png)

```go
func main() {
   /*
      输入和输出：fmt包

      输出：
         fmt.Println()  //打印换行
         fmt.Print()     //打印
         fmt.Printf()   //格式化打印
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
```




### :trophy:编码规范

---

#### :four_leaf_clover:为什么需要代码规范

+ 代码规范不是强制的，也就是不遵循代码规范写出来的代码也是完全没有问题的
+ 代码规范的目的是方便团队形成统一的代码风格，提高代码的可读性，规范性和统一性，本规范将从命名规范，注释规范，代码风格和Go语言提供常用工具这几方面左一个说明。
+ 规范并不是唯一的，也就是说理论上每个公司都可以制定自己的规范，不过一般来说代码的整体规范差异性不是很大。

#### :four_leaf_clover:代码规范

> 命名规范

命名是代码中很重要的一部分，统一的命名规则有利于提高代码的可读性，好的命名仅仅通过，欧美命名就可以获取到很多信息。

+ 当命名（包括常量、变量、类型、函数名、结构字段等）以一个大写字母开头。如：Group1，那么使用这种形式的标识符的对象就可以被外部的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的private）
+ 命名如果以小写字母开头，则对包外事不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的private）

**包名：package**

保持package名字和目录名保持一致，尽量采取有意义的包名，简短，有意义，尽量不要和标准库冲突，包名应该为小写单词，不要使用下划线或者大小写混写。

```go
package model
package main
```

**文件名**

尽量采用有意义的文件名，简短，有意义，应该为小写单词，使用下划线分隔各个单词（蛇形命名）

```go
user_model.go
```

**结构体命名**

+ 采用驼峰法，首字母根据访问控制大写或小写
+ struct 申明和初始化格式采用多行，例如下面：

```GO
// 多行声明
type User struct {
	Username string
	Email    string
}

func main() {
    //多行初始化
	u := User{
		Username: "baby",
		Email:    "baby@123.com",
	}
	fmt.Println(u)
}
```

**接口命名**

+ 命名规则基本上和上面的结构体类似
+ 单个函数的结构体名以‘er’作为后缀，例如 Reader、Writer

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

**变量命名**

+ 和结构体类似，变量命名一般遵循驼峰法，首字母根据访问控制原则大写或小写，单遇到特有名词时，需要遵循以下规则：
+ 如果变量为私有，且特有名词为首个单词，则使用小写，如apiClient
+ 其他情况都应使用该名词原有的方法，如APIClient、repoID、UserID
+ 错误示例：UrlArray，应该写成 urlArray 或者 URLArray

```GO
var isExist bool
var hasCongict bool
var canManage bool
var allowGitHook bool
```

**常量命名**

常量均需要使用全部大写字母组成，并使用下划线分词

```
const APP_VAR = "1.0"
```

如果事枚举类型的常量，需要首先创建相应类型：

```GO
type Scheme string

const (
	HTTP Schem = "http"
    HTTPS Schem = "https"
)
```



#### :four_leaf_clover:注释

Go提供C风格的/**/块注释，如C++风格的 //注释，行数值时常态，块注释主要显示为包注释，单在表达式中很有用，或者禁用大量代码。

+ 单行注释：是常见的注释形式，可以在任何地方使用//注单行释
+ 多行注释：也叫块注释，均已`/*`开头，并以`*/`结尾，不可以且套使用，多行注释一般在文档描述或注释成块的代码片段，Go语言自带的fofoc工具可以根据注释生成文档，生成自动对应的网站，（dolang.org就是用godoc直接生成的），注释质量决定了生成的文档质量，每个包都应该有以恶搞包注释，在package子句前有一个块注释，对于多文件包，包注释只需要存在于一个文件中，任何一个都可以。包评论应该介绍包，并提供整个包相关信息，它将首先出现在godoc页面上，并设置下面的详细文档。

**包注释**

每个包都应该有一个包注释，一个位于package子句前的块注释或者行注释，如果包有多个文件，只需要任何一个go文件中（一般和包同名的文件）即可。包注释应该包含下面几个基本信息（请严格按照这个顺序，简介，创建人，创建时间）

+ 包的基本简介（包名，简介）
+ 创建者，格式：创建人：人名
+ 创建时间，格式：创建时间：yyyyMMdd

```go
/*
util包，该报包含了项目共用的一些常量，封装了项目中的一些共用函数
创建人：zhangsan
创建时间：2099.11.13
/*
```

**结构体（接口）注释**

每个自定义的结构体或者接口都应该右注释说明，该结构体进行简要注释，放在结构体的前一行，格式为：结构体名，结构体说明，同时结构体内的每个成员变量都要有说明，该说明放在成员变量后面（注意对齐），实例如下

```GO
 //user.用户对象，定义了用户信息
type User struct {
	Username string	//用户名
	Email    string	//邮箱
}
```

**函数（方法）注释**

每个函数，或者方法（结构体或者接口下的函数称为方法）都应该右注释说明，函数的注释应该包括三个方面（严格按照此顺序撰写）：

+ 简要说明，格式说明：一函数名开头，“，”分割说明部分
+ 参数列表：每行一个参数，参数名开头，“，”分割说明
+ 返回值：每行一个返回值

```GO
//NewtAttrModel,属性数据层操作类的工厂方法
//参数，
//		ctx,上下文信息
//返回值，
//		属性操作类指针
func NewAttroModel(ctx *common.Context) *AttrModel {
    
}
```

代码逻辑注释

对于一些关键位置的代码逻辑，或者局部较为复杂的逻辑，需要有相应的逻辑说明，方便其他开发者阅读该段代码，实例如下：

```GO
// 从Redis中批量读取属性，对于没有读取到的id，记录到一个数组里面，准备从DB中读取
xxxxx
xxxxx
xxxxx
```

**注释风格**

统一使用中文注释，对于中英文字符之间严格使用空格分隔， 这个不仅仅是中文和英文之间，英文和中文标点之间也都要使用空格分隔，例如：

```GO
// 从 Redis 中批量读取属性，对于没有读取到的 id ， 记录到一个数组里面，准备从 DB 中读取
```

上面 Redis 、 id 、 DB 和其他中文字符之间都是用了空格分隔。

- 建议全部使用单行注释
- 和代码的规范一样，单行注释不要过长，禁止超过 120 字符。



#### :four_leaf_clover: import规范

import在多行的情况下，goimports会自动格式化，但是还时需要规范一下的import的一些规范，如果你在一个文件里面引入了一个pacage，换徐建议采用如下格式：

```GO
import (
	"bufio"
	"fmt"
	"os"
)
```

如果你的报引用三种类型的报、标准库报、程序内部包、第三方包，建议：

```GO
import (
	"bufio"
	"fmt"
	"os"
    
    "github.com/xxx"
    "github.com/yyy"
)
```

有顺序的引入包，不同类型采用空格分离，第一种是标准库、第二种项目包，第三种第三方包。

在项目不要使用相对路径引入包

```GO
// 这是不好的导入
import "../net"

//这是正确的做法
import ”github.com/repo/proj/src/net“
```

但是如果引入本地项目的其他包，最好使用相对路径



#### :four_leaf_clover: 错误处理

- 错误处理的原则就是不能丢弃任何有返回err的调用，不要使用 _ 丢弃，必须全部处理。接收到错误，要么返回err，或者使用log记录下来
- 尽早return：一旦有错误发生，马上返回
- 尽量不要使用panic，除非你知道你在做什么
- 错误描述如果是英文必须为小写，不需要标点结尾
- 采用独立的错误流进行处理



```go
// 错误写法
if err != nil {
  // error handling
} else {
  // normal code
}

// 正确写法
if err != nil {
  // error handling
  return // or continue, etc.
}
// normal code
```



#### :four_leaf_clover: 测试

单元测试文件名命名规范为 example_test.go 测试用例的函数名称必须以 Test 开头，例如：TestExample 每个重要的函数都要首先编写测试用例，测试用例和正规代码一起提交方便进行回归测试





### :trophy:流程控制

---

程序的流程控制结构一共有三种：顺序结构，选择结构，循环结构

顺序结构：从上到下，逐行执行，默认的逻辑

选择结构：条件满足某些代码才会执行

+ if
+ switch
+ select
+ channel

循环结构：条件满足某些代码会被反复执行0-N次

+ for



#### :four_leaf_clover: if 语句

条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为true来决定是否执行指定语句，并在条件为false的情况下执行另外的语句。

下图展示了程序语言中条件语句的结构：

  ![image-20221108221307856](/images/if.png)

```go
func main() {

   //分数
   var score int = 90

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
   } else {                        //以上条件都不满足就走else
      fmt.Println("不及格")
   }
}
```

> if 嵌套语句

```sh
if 布尔表达式 1 {
	// 在布尔表达式 1 为 true 时执行 
	if 布尔表达式 2 {
		// 在布尔表达式 2 为 true 时执行 
	}
}
```


示例：

```go
// if 语句嵌套
func main() {

   var a int = 100
   var b int = 200

   if a == 100 {
      fmt.Println("a 满足条件")
      if b == 200 {
         fmt.Println("b 满足条件")
      }
   }
}
```

验证密码案例：（再次输入密码）

```GO
// if 嵌套语句验证密码
func main() {

   // 验证密码，再次输入密码
   var a, b string
   var pwd string = "go@123"
   // 用户输入
   fmt.Print("请输入密码:")
   fmt.Scanln(&a)
    // 业务：验证密码是否正确
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
```





#### :four_leaf_clover: switch

对于if {} else if {} 写的多了会让代码很多

switch语句用于基于不同条件执行不同的动作，每一个case分支都是唯一的，从上至下注意测试，直到匹配为止；

```GO
	switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
	}
```

switch语句执行的过程从上至下，匹配找到的匹配项，匹配项后面也不需要再加break，switch默认情况下case最后自带break语句。

```GO
// switch语句
func main() {
	var score int = 90

	// 匹配case
	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 50, 60, 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
```



switch省略后面的变量，详单与直接作用在true上，相当于默认为 switch true{}

```go
// switch 默认的条件 bool = true
switch {
case false:
	fmt.Println("false")
case true:
	fmt.Println("true")
default:
	fmt.Println("其他")
}
```

> fallthrough 贯穿；直通

开发很少使用，了解

switch默认情况下匹配成功后就不会执行去他的case，如果我们需要执行后面的case，可以将fallthrough穿透case，使用fallthrough会强制执行后面的case语句，fallthrough不会判断下一条case的表达式结果是否为true；

```go
func main() {
   a := false
   switch a {
   case false:
      fmt.Println("1.case的条件为false")
      fallthrough // case穿透，不管下一条条件是否满足，都会执行
   case true:
      if a == false {
         break  //终止穿透
      }
      fmt.Println("2.case的条件为true")
   }
}
```



#### :four_leaf_clover: for

不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复性执行某些语句。

for循环是一个循环控制结构，可以执行指定次数的循环。

```gO
func main() {
	// for 给控制变量赋值，循环条件，给控制变量增量或减量
	for i := 1; i <= 5; i++ {
		// 循环体
		fmt.Println("打印次数:", i)
	}
}
```

计算1到10的和

```GO
func main() {
   // 计算1到10的和
   sum := 0
   for i := 1; i <= 10; i++ {
      sum = sum + i
   }
   fmt.Println(sum)
}
```

for 可以拆分开写

三种：

+ for 给控制变量赋值;循环条件;给控制变量增量或减量{}
+ 循环条件;给控制变量增量或减量{}
+ for{}

```GO
i := 1
for {
   fmt.Println(i)
   i++
   if i == 15 {
      break
   }
}
```

> 案例

**案例一：for 打印方阵**

```go
/* 打印一个方阵
* * * * * * 换行
* * * * * *
* * * * * *
* * * * * *
* * * * * *
 */
func main() {

   //for i := 0; i < 5; i++ {
   // fmt.Print("* ")
   //}

   for j := 0; j < 5; j++ {
      for i := 0; i < 5; i++ {
         fmt.Print("* ")
      }
      fmt.Println()
   }
}
```

**案例二：打印九九乘法表**

拆分思路：（重要）

```go
// 打印99乘法表
func main() {
   fmt.Println("-------------Go打印九九乘法表-------------")
   for i := 1; i < 10; i++ {
      for j := 1; j <= i; j++ {
         fmt.Printf("%dx%d=%d\t", i, j, i*j)
      }
      fmt.Println()
   }
}
```

退出for循环，两个关键字

1. break：结束整个循环
2. continue：结束当此循环

```go
func main() {
   // break 结束当前循环
   for i := 0; i < 10; i++ {
      if i == 5 {
         break
      }
      fmt.Print(i)
   }
   fmt.Println()

   // continue 结束本次循环
   for i := 0; i < 10; i++ {
      if i == 5 {
         continue
      }
      fmt.Print(i)
   }
}
```



#### :four_leaf_clover:String

> 什么是tring

Go中的字符串是一个字节的切片，可以通过将其内容封装在`“”`中创建字符串，Go中的字符串是Unicode兼容的，并且是UTF-8编码，字符串是一些字符的集合。

```go
func main() {
   str := "www.baidu.com"
   fmt.Println(str)

   // 获取字符串的长度 len
   fmt.Println("字符串的长度为:", len(str))

   // 获取指定的字节，默认从0开始
   //acsii 编码 119
   fmt.Println("字节打印:", str[0])
   fmt.Printf("%c", str[0])

   // for 遍历
   for i := 0; i < len(str); i++ {
      //fmt.Println(str[i])
      fmt.Printf("%c", str[i]) //wwww.baidu.com
   }

   // for range循环，便利数组、切片...
   // 返回下表和值
   for i, v := range str {
      fmt.Print(i)
      fmt.Printf("%c", v)
   }

   // string 不能修改的
   // cannot assign to str[2] (value of type byte)
   str[2] = 'A'
   fmt.Println(str[2])
}
```



### :trophy:函数

---

+ 函数是基本的代码块，用于执行一个任务
+ Go语言最少有个main函数
+ 可以通过函数来划分不同的功能，逻辑上每个函数执行的是指定的任务，
+ 函数声明告诉了编译器函数的名称，返回类型和参数

```
// main 函数
func main() {
   fmt.Println("hello,world")

   fmt.Println(add(1, 2))
}

/*
   func 函数名(参数,参数...) 函数调用后的返回值 {
      函数体   //执行一段代码
      reture //返回结果
   }
*/
func add(a, b int) int {
   c := a + b
   return c
}
```

#### :four_leaf_clover:函数的声明

Go语言函数定义格式如下：

```go
func function_name( [parameter list]) [return_type] {
    函数体
}
```

+ 无参无返回值函数
+ 有一个参数的函数
+ 有两个参数的函数
+ 有一个返回值的函数
+ 有多个返回值的函数

```go
package main

import "fmt"

/*
+ 无参无返回值函数
+ 有一个参数的函数
+ 有两个参数的函数
+ 有一个返回值的函数
+ 有多个返回值的函数
*/
func main() {
   // 函数调用
   printinfo()
   myprint("有一个参数的函数")
   myprint(addstring("有两个参数的函数\t", "有一个返回值的函数"))
   x, y := swap("one", "two")
   fmt.Println(x, y)
}

// 无参无返回值函数
func printinfo() {
   fmt.Println("无参无返回值函数")
}

// 有一个参数的函数
func myprint(msg string) {
   fmt.Println(msg)
}

// 有两个参数的函数
func addstring(a, b string) string {
   c := a + b
   return c
}

// 有多个返回值的函数
func swap(x, y string) (string, string) {
   return y, x
}
```

比较两个数字大小

```go
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
```



#### :four_leaf_clover:可变参数

概念：一个函数的参数类型确定，但个数不确定，就可以使用可变参数。

```GO
func myfunc(arg ..int){}
```

`arg  ...int` 告诉 Go 这个函数接受不定数量的参数，类型全部是 int

```go
func main() {
   getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

// ... 可变参数
func getSum(nums ...int) {
   sum := 0

   for i := 0; i < len(nums); i++ {
      fmt.Println(nums[i])
      sum += nums[i]
   }
   fmt.Println("sum:", sum)
}
```

注意事项：

+ 如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在列表的最后
+ 一个函数的参数列表中最多只能有一个可变参数

#### :four_leaf_clover:参数传递

按照数据的存储特点来分：

+ 值类型的数据：操作的是数据本身、int、string、bool、float64、array...
+ 引用类型的数据：操作的是数据的地址slice、map、chan

函数中的重点：

> 值传递

```go
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
```

> 引用传递

操作的始终是同一个空间

**变量在内存中是存放在一定的地址上的，修改了变量实际是修改了变量地址处的内存**

```go
// 引用传递
func main() {

	// 切片，可以扩容的数组
	s1 := []int{1, 2, 3, 4}
	fmt.Println("默认的数据:", s1) // [1 2 3 4]
	// 传入的是引用类型的数据，地址
	update2(s1)
	fmt.Println("调用后的数据:", s1) // [100 2 3 4]
}
func update2(s2 []int) {
	fmt.Println("传递的数据:", s2) // [1 2 3 4]
	s2[0] = 100
	fmt.Println("修改后的数据:", s2) // [100 2 3 4]
}

```

#### :four_leaf_clover:函数中变量的作用域

作用域：变量可以使用的范围

局部变量：函数内部定义的变量：局部变量

全部变量：函数外部定义的变量：全局变量

注意：无论是全局变量还是局部变量都遵循就近原则！

```go
// 全局变量
var num int = 100

func main() {
   // 函数体内的局部变量
   temp := 100
   // if、for语句定义的一次性变量局部变量
   if b := 1; b <= 10 {
      // 语句内的局部变量
      temp := 50
      fmt.Println(temp) // 局部变量，就近原则
      fmt.Println(b)
   }
   fmt.Println(temp)
   fmt.Println(num)

   f1()
   f2()
}

// 内部
func f1() {
   a := 1
   fmt.Println(a)
   num := 30 // 就近原则
   fmt.Println(num)
}
func f2() {
   //fmt.Println(a)   // 不能使用其他函数定义的变量
   fmt.Println(num)
}
```



#### :four_leaf_clover:递归函数

定义：一个函数自己调用自己，就叫递归函数

注意：递归函数需要一个出口，逐渐向出口靠近，没有出口就会形成死循环。

求和：1,2,3,4,5

```GO
getsum(5)
	getsum(4)+5
		getsum(3)+4
			getsum(2)+3
				getsum(1)+2
					1
```

递归十分消耗程序的内存

```GO
func main() {
	sum := getSum(5)
	fmt.Println(sum)
}

// 5				5
// getSum(4) + 5	15
// getSum(3) + 4	10
// getSum(2) + 3	6	getSum(2) = getSum(1) + 2
// getSum(1) + 2	3	== 1 + 2
// 1				1
func getSum(n int) int {
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}
```



#### :four_leaf_clover:defer

defer语义：推迟、延迟

在go语言中，使用defer关键字来推迟一个韩硕或者方法的执行

```go
func main() {
	f("1")
	fmt.Println("2")
	defer f("3") // 会被延迟到最后执行
	fmt.Println("4")
	defer f("5")
	defer f("6") //defer语句会按照逆序执行
}

func f(s string) {
	fmt.Println(s)
}
```



defer韩硕或者方法：一个函数或方法的执行被延迟了

+ 你可以子函数中添加多个defer语句，当函数执行到最后时，这些**defer语句会按照逆序执行**，最后该函数返回，特别是当你在进行进行一些打开资源时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。
+ 如果有很多调用defer，那么**defer是采用后进先出（栈）模式**

**值是什么时候传递进去的:**

常用于关闭文件操作，io、线程、网络

```go
func main() {
	a := 10
	fmt.Println("a=", a)
	defer f3(a) // 函数中的a= 10;参数已经传递进入，在最后执行
	a++
	fmt.Println("a=", a)

}
func f3(s int) {
	fmt.Println("函数中的a=", s)
}
```



#### :four_leaf_clover:高级：函数的数据类型

函数的数据类型：func(参数类型)（返回值类型）

函数也是一种数据类型

```GO

// func() 本身就是一个数据类型；那么就可以定义一个函数类型的变量
func main() {
	// 不加括号，函数就是一个变量
	// f0()  如果加上了括号那就变成了函数
	fmt.Printf("%T\n", f0)  // func() | func(int, int) | func(int, int) int
	fmt.Printf("%T\n", 10)  // int
	fmt.Printf("%T\n", "A") // string

	// 定义函数类型的变量
	var f5 func(int, int)
	f5 = f0
	// 调用 f5
	f5(1, 2)
	fmt.Println(f5) // 0xbdfe40
	fmt.Println(f0) // 0xbdfe40

}
func f0(a, b int) {
	fmt.Println(a, b)
}
```

func() 本身就是一个数据类型；不加括号就是一个变量，变量可以赋值，赋值给一个对应的数据类型函数

**函数在Go语言中是符合函数类型，可以看做是一种特殊的变量**



#### :four_leaf_clover:高级：匿名函数

Go语义是支持函数式编程：

1、将匿名函数作为另外一个函数的参数，回调函数

2、降匿名函数作为另外一个函数的返回值，可以形成闭包结构

```go
// 匿名函数
func main() {
   f6()
   f7 := f6 // 函数本身也是一个变量
   f7()

   // 匿名函数
   f8 := func() {
      fmt.Println("f8函数")
   }
   f8()

   //简化；匿名哈桉树自己调用自己，只执行一次
   func() {
      fmt.Println("f9函数")
   }()

   //匿名函数传参
   func(a, b int) {
      fmt.Println(a, b)
   }(1, 2)

   //匿名函数传参 + return
   r1 := func(a, b int) int {
      fmt.Println(a, b)
      return a + b
   }(1, 2)
   fmt.Println(r1)
}

func f6() {
   fmt.Println("f6函数")
}
```



#### :four_leaf_clover:高级：回调函数

高价函数：根据go语言的数据特点，**可以将一个函数作为另外一个函数的参数**。

func1()，func2()

将func1函数作为func2这个函数的参数

+ func2函数：叫做高阶函数，接受了一个函数作为参数的函数
+ fucn1函数：叫做回调函数，作为另外一个函数的参数

```go
func main() {
   r1 := add1(1, 2)
   fmt.Println(r1)

   r2 := oper(3, 4, add1)
   fmt.Println(r2)

   r3 := oper(10, 4, add1)
   fmt.Println(r3)

   // 既然可以调用实际函数，那么也可以调用匿名函数
   r4 := oper(12, 4, func(a int, b int) int {
      if b == 0 {
         fmt.Println("除数不能为0")
         return 0
      }
      return a / b
   })
   fmt.Println(r4)
}

// 高阶函数,可以接受一个函数作为参数
func oper(a, b int, fun func(int, int) int) int {
   r := fun(a, b)
   return r
}

func add1(a, b int) int {
   return a + b
}
func sub(a, b int) int {
   return a - b
}
```

在遇到：像使用同一个方法实现不同的功能，就可以使用回调函数实现



#### :four_leaf_clover:高级：闭包

一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量

并且该外层函数的返回值就是这个内层函数。

这个内层函数的和外层函数的局部变量，统称为闭包结构

局部变量的生周期就会发生改变，正常的局部变量会随着函数的调用而创建，随着函数的结束而销毁；

但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还在继续使用



注意：开发中避免这样使用，阅读起来会产生错觉

```go
package main

import "fmt"

func main() {

   r1 := increment()
   fmt.Println(r1) // 地址：0x9de7e0

   v1 := r1()
   fmt.Println(v1) // 1
   v2 := r1()
   fmt.Println(v2)   // 2
   fmt.Println(r1()) // 3
   fmt.Println(r1()) // 4
   fmt.Println(r1()) // 5

   r2 := increment()
   fmt.Println(r2)   // 地址：0x9de7c0
   fmt.Println(r2()) // 1

}

// 自增
func increment() func() int {
   // 局部变量 i
   i := 0
   // 定义一个匿名函数，给变量自增并返回
   fun := func() int { // 内层函数，没有执行的
      i++
      return i
   }
   return fun
}
```
















