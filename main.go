package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 2
	b := isHappy(num)
	fmt.Println(b)
}

// 这种方法会造成内存溢出
func isHappy(n int) bool {
	numStr := strconv.Itoa(n)
	//查看类型
	//fmt.Printf("numStr 类型 %T \n", numStr)

	//申明一个切片
	numArr := make([]int, 0)
	for _, s := range numStr {
		//fmt.Printf("s 类型 %T \n", s)
		s2 := string(s)
		atoi, _ := strconv.Atoi(s2)
		numArr = append(numArr, atoi)
	}
	sum := 0
	for _, i := range numArr {
		sum += i * i
	}
	if sum == 1 {
		return true
	} else {
		if sum == n {
			return false
		}
		return isHappy(sum)
	}
}
