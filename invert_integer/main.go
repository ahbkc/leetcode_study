package main

import (
	"fmt"
	"strconv"
)

/*
题目：7.给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2的31次方,  2的31次方 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
*/

//结果：测试通过，已提交

func main() {
	for i := 0; i < 1; i++ {
		r := 1534236469
		fmt.Println("传入值：", r)
		fmt.Println("返回值: ", reverse(r))
	}
}

func reverse(x int) int {
	var max = 2<<31 - 1
	array := []rune(strconv.Itoa(x))
	var s string
	var length = len(array) - 1
	for i := length; i >= 0; i-- {
		s += string(array[i])
	}
	if x < 0 {
		s = "-" + s[:len(s)-1]
	}
	n, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	if n >= max || n <= -max {
		return 0
	}
	return n
}
