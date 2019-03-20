package main

import (
	"fmt"
	"testing"
)

type Data struct {
	Left   []int
	Right  []int
	Result string
}

var (
	TestData = []Data{
		{
			Left:   []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			Right:  []int{5, 6, 4},
			Result: "1000000000000000000000000000466",
		},
		{
			Left:   []int{0},
			Right:  []int{7, 3},
			Result: "37",
		},
		{
			Left:   []int{0},
			Right:  []int{0},
			Result: "0",
		},
		{
			Left:   []int{1, 8},
			Right:  []int{0},
			Result: "81",
		},
		{
			Left:   []int{9, 8},
			Right:  []int{1},
			Result: "90",
		},
	}
)

func TestAddTwoNumber(t *testing.T) {
	for _, v := range TestData {
		var l1, l2 = getLinkedList(v.Left), getLinkedList(v.Right)
		l3 := addNum(&l1, &l2)
		var result string
		for l3 != nil {
			result = fmt.Sprint(l3.Val) + result
			t.Logf("%s\n", result)
			l3 = l3.Next
		}
		if v.Result != result {
			t.Errorf("计算失败, v.Result: %s, result: %s\n", v.Result, result)
		}
	}
}

// 获取链表结构体
func getLinkedList(data []int) ListNode2 {
	if data == nil || len(data) == 0 {
		return ListNode2{}
	}
	var n1 = ListNode2{}
	var temp1, temp2, temp3 *ListNode2 = &n1, nil, nil
	for k, v := range data {
		temp2 = temp1 // 赋值给临时变量
		temp2.Val = v // 赋值 Val 字段
		if k == len(data)-1 {
			continue
		}
		temp3 = temp1               // n1.Next  将指针保存到临时变量, 并清空temp1变量所指向的值
		temp3.Next = new(ListNode2) // 为下一个节点审请空间
		temp1 = temp3.Next
	}
	return n1
}
