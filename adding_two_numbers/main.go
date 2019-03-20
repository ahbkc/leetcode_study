package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

/*
2. 两数相加
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

// 失败

func main() {
	var first = &ListNode{Next: &ListNode{}}
	var second = &ListNode{Next: &ListNode{}}
	var firstArray = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	a := int64(1000000000000000000000000000001)
	fmt.Println(a)
	var secondArray = []int{5, 6, 4}
	var tempNode1, tempNode2 = first.Next, second.Next
	for i := 0; i < len(firstArray); i++ {
		if tempNode1.Next == nil {
			tempNode1.Val = firstArray[i]
			tempNode1.Next = &ListNode{}
		}
		tempNode1 = tempNode1.Next
	}
	for i := 0; i < len(secondArray); i++ {
		if tempNode2.Next == nil {
			tempNode2.Val = secondArray[i]
			tempNode2.Next = &ListNode{}
		}
		tempNode2 = tempNode2.Next
	}
	result := addTwoNumbers(first, second)
	fmt.Println(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var temp1 = l1.Next
	var temp2 = l2.Next
	var v1, v2 = fmt.Sprint(l1.Val), fmt.Sprint(l2.Val)
	for temp1 != nil {
		v1 = fmt.Sprint(temp1.Val) + v1
		temp1 = temp1.Next
	}
	for temp2 != nil {
		v2 = fmt.Sprint(temp2.Val) + v2
		temp2 = temp2.Next
	}
	val1, _ := strconv.Atoi(v1)
	val2, _ := strconv.Atoi(v2)
	fmt.Println(unsafe.Sizeof(val1))
	if val2 == 0 || val2 < 0 {
		return l1
	}
	if val1 == 0 || val1 < 0 {
		return l2
	}
	//fmt.Printf("l1 的值是: %s, l2的值是: %s, 相加结果是: %d\n", v1, v2, val1+val2)
	//分割结果
	var vv = val1 + val2
	result := fmt.Sprint(vv)
	valArray := strings.Split(result, "")
	var array []*ListNode
	for _, v := range valArray {
		val, _ := strconv.Atoi(v)
		array = append(array, &ListNode{Val: val})
	}
	for i := len(array) - 1; i >= 0; i-- {
		if i-1 < 0 {
			continue
		}
		array[i].Next = array[i-1]
	}
	resultNode := array[len(array)-1:]
	return resultNode[0]
}
