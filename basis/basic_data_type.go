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

}
