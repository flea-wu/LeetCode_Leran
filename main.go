package main

import (
	"LeetCode_Leran/algorithm"
	"LeetCode_Leran/utils"
	"fmt"
)

func main() {
	var i = 19
	algorithm.IsHappy(i)
	utils.SayHello()
}

func init() {
	fmt.Println("main 包下的 init 初始化函数")
}
