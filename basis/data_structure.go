package basis

import (
	"fmt"
	"unsafe"
)

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

	// 浮点类型 float32 float64
	var float32Var float32 = 1
	fmt.Println(float32Var)
	fmt.Printf("类型%T，占用空间为：%d", float32Var, unsafe.Sizeof(float32Var))

}
