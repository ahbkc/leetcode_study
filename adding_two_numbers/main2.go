package main

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

// add-two-numbers  2019/3/20 通过

func addNum(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	var n1, n2 = l1.Next, l2.Next
	if n1 == nil && n2 == nil {
		var sum int
		if sum = l1.Val + l2.Val; sum > 9 {
			return &ListNode2{Val: sum - 10, Next: &ListNode2{Val: sum / 10}}
		}
		return &ListNode2{Val: sum}
	}

	// 不允许使用字符串转整型值
	node := new(ListNode2)
	var temp1, temp2, temp3 *ListNode2 = node, nil, nil
	var overflowValue int
	// input:  [1,2] [9,1]
	// output: [0,4]
	n1, n2 = l1, l2
	for n1 != nil || n2 != nil || overflowValue == 1 {
		temp2 = temp1
		var v1, v2 int
		if n1 == nil {
			v1 = 0
		} else {
			v1 = n1.Val
		}
		if n2 == nil {
			v2 = 0
		} else {
			v2 = n2.Val
		}
		var sum = v1 + v2
		if overflowValue > 0 {
			sum += overflowValue
			overflowValue = 0 // 清空
		}
		if sum > 9 {
			temp2.Val = sum - 10
			overflowValue = 1
		} else {
			temp2.Val = sum
		}
		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}

		if n1 == nil && n2 == nil && overflowValue == 0 {
			continue
		}

		temp3 = temp1
		temp3.Next = new(ListNode2)
		temp1 = temp3.Next
	}
	return node
}
