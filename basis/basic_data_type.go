package basis

import (
	"fmt"
	"strconv"
	"unsafe"
)

// golang 整数类型分: 有符号和无符号，int 和 uint 和系统优化
func DataStructure() {
	// 整数类型 int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 byte
	// byte 和 uint8 等价 当要存储字符时选用 byte
	var hello int8 = 1
	fmt.Println("hello=", hello)

	var uintName = -1
	fmt.Println(uintName)

	varName := "hello"
	fmt.Println(varName)

	// 格式化输出
	fmt.Printf("类型%T，占用空间为：%d", varName, unsafe.Sizeof(varName))
	fmt.Println()

	// 浮点类型 float32 float64，浮点数尾数部分可能造成精度丢失；float64 的精度要比 float32 要高；
	// 浮点类型存储内容： 符号位 + 指数位 + 尾数位，在存储过程中，精度会丢失
	// 浮点类型长度固定，不受操作系统影响
	//
	var float32Var float32 = 1
	fmt.Println(float32Var)
	fmt.Printf("类型%T，占用空间为：%d", float32Var, unsafe.Sizeof(float32Var))

	fmt.Println()
	float64Var := 8.9999999
	fmt.Println(float64Var)
	fmt.Printf("类型%T，占用空间为：%d", float64Var, unsafe.Sizeof(float64Var))
	fmt.Println()

	// 字符类型 byte
	// go中 字符常量通常是用单引号引起来的单个字符 'f'
	var byteVar = 'f'
	fmt.Println(byteVar)
	// 字符类型输出
	fmt.Printf("byteVar的字符类型%c  数字类型%d", byteVar, byteVar)
	printf, _ := fmt.Printf("%c", byteVar)
	fmt.Println()
	fmt.Println(printf)
	fmt.Println()

	// 布尔类型 bool 只允许取 true 和 false；bool 类型只占一个字节
	var boolVar = true
	fmt.Println(boolVar)
	fmt.Println()

	// string（字符串） 类型，go 中字符串是由单个字节连接起来的，一旦赋值就不能修改
	// 字符串两种表现形式 1.双引号，会识别转义字符 2.反引号，以字符串原生形式输出，包括换行和转义
	// 字符串拼接使用 +

	stringVar := "我是字符串"
	fmt.Println(stringVar)

	stringVar2 := `解决解决军\n \t`
	fmt.Println(stringVar2)

	// 基本类型默认值
	// 整型 0，浮点型 0，字符串 "" ，布尔 false

	// 基本数据类型转换
	// go 中数据类型都需要显示类型转换，不能自动转换
	// 表达式为：T(v) ，即将值 v 转换为类型 T
	var intVar88 = 99
	i := int8(intVar88)
	fmt.Println(i)

	// 基本数据类型和 string 类型转换
	// 可以直接使用 Sprintf 进行转换
	sprintf := fmt.Sprintf("%d", 1)
	fmt.Println(sprintf)
	// 使用 strconv 包进行转换
	strconv.Itoa(1)
	strconv.FormatBool(true)

	// string 转基本数据类型；如果没转换成功，转化后的值可能为 默认值
	parseBool, _ := strconv.ParseBool("true")
	fmt.Println(parseBool)

	// 指针类型
	a := 5
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	// 值类型和引用类型
	// 值类型：基本数据类型int系列，float系列，bool，string，数组，结构体struct
	// 引用类型：指针，slice切片，map，管道chan，interface，等都是引用类型
	// 值类型：变量直接存储值，通常存储在内存栈中
	// 引用类型：变量存储地址

	if 1 > 2 || 3 < 4 {

	}

	// 自定义类型

	for a := 1; a < 10; a++ {
		fmt.Println("hhhhh")
	}

	// 函数也是一种数据类型，func

	// go 支持自定义类型
	// type 自定义类型名称 数据类型；相当于一个别名

	type helloType int
	var helloVar helloType = 1
	fmt.Printf("%T", helloVar)
	fmt.Println()

	// init 函数，每个原文件都可以有一个 init 函数
	// 在 main.go 中 全局变量 > init 函数 > main 函数
	// 在 其他.go 和 main.go 中 其他.go 中的 全局变量 > init > main.go 中 全局变量 > init 函数 > main 函数

	// 匿名函数
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(1, 3)
	fmt.Println("结果为：", res)

	// 将匿名函数赋值给变量
	res2 := func(n1 int, n2 int) int {
		return n1 + n2
	}

	i2 := res2(1, 2)
	fmt.Println(i2)
	// 全局匿名函数，将匿名函数赋值给全局变量

	// 闭包，闭包是类，函数是操作， n是字段；函数和它使用的字段构成闭包

	// defer 函数，当函数执行完，再执行 defer 栈中函数

	defer func() {
		fmt.Println("defer...")
	}()

	fmt.Println("defer end")

	// defer 将语句放入到栈中时，也会将相关的值拷贝进栈

	// defer 主要作用用于关闭资源

	// 函数的传参方式
	// 1. 值传递，2. 引用传递
	// 不管是值传递还是引用传递，都是拷贝，值传递拷贝的是值，引用传递拷贝的是地址

	// 操作字符串常见函数
	// 遍历字符串，需要将字符串转换成 []rune

	// 时间和日期相关函数 time

	// go 一些内置函数 len 求长度；new 主要用来值类型分配内存，返回的是指针 例：new(int)
	// make 也是来分配内存，主要用来分配引用类型的内存，例如：channel，map，slice

	// go 中的异常
	// panic 相当于java中的 throw，recover 相当于 try-catch
	// 自定义错误 error.New(xxx)

	// 数组和切片
	var arr [7]int
	arr[0] = 1
	i3 := &arr
	fmt.Println(i3)
	// 数组地址等于数组元素第一个，for-range 遍历数组

	// 切片，1.对数组的一个引用 2.make
	// 基本语法
	var myslice []int = make([]int, 1, 1)
	// cap 容量（可选，cap>=len），len 大小
	fmt.Println(myslice)

}
