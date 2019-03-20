package main

import "fmt"

/*
题目：1.两数之和

给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

//output: [0, 1]
//leetcode上完成的第一个算法题, 20180921 14:27

func main() {
	var nums = []int{2, 5, 10}
	var target = 7
	if array := twoSum(nums, target); len(array) > 0 {
		fmt.Println("target number: ", target)
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i])
		}
	}
}

//相同下标值不同被重复利用
func twoSum(nums []int, target int) []int {
	var array []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				array = append(array, i)
				array = append(array, j)
				return array
			}
		}
	}
	return array
}
