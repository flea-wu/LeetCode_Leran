package algorithm

import (
	"fmt"
	"strconv"
)

// 这种方法会造成内存溢出
func IsHappy(n int) bool {
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
		return IsHappy(sum)
	}
}

func IsHappyNum01(num int) bool {

	set := map[int]bool{}
	//  num, set[num] = step(num), true 代码解释: step(num) 返回值赋值给num，true 赋值给 set[num]
	for ; num != 1 && !set[num]; num, set[num] = step(num), true {
	}
	//
	for {
		if num != 1 && !set[num] {
			num := step(num)
			set[num] = true
		}
		return num == 1
	}
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

//init函数不需要传入参数也没有返回值，而且init函数是不能被其他函数调用的。
//在一个文件中也可以有多个init函数
//只是为了执行init函数而导入包：import _ "image/png"
func init() {
	fmt.Println("algorithm 包下面的初始化函数1")
}

func init() {
	fmt.Println("algorithm 包下面的初始化函数2")
}
