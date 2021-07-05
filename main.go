package main

import (
	"LeetCode_Leran/basis"
	"container/list"
	"fmt"
	_ "math"
	"runtime"
)

func main() {
	basis.DataStructure()
}

//func longestCommonPrefix(strs []string) {
//	datamap := make(map[string]string)
//	l := list.New()
//	for _, str := range strs {
//		l.PushFront(len(str))
//
//
//		split := strings.Split(str, "")
//		//for _, char := range split {
//		//
//		//}
//	}
//
//}

func test() {
	l := list.New()
	l.PushFront("a")

	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()

	//类似 java 中，主动抛出一个异常
	panic("crash")

	fmt.Println("==========")
	fmt.Println(&l)
}

func init() {
	fmt.Println("main 包下的 init 初始化函数")
}
