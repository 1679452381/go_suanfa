package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

// 遍历链表
func GetListNodeElements(head *ListNode) []int {
	cur := head
	listNodeValues := make([]int, 0)
	for cur != nil {
		listNodeValues = append(listNodeValues, cur.Value)
		cur = cur.Next
	}
	return listNodeValues
}

// 双指针法
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode // 空指针

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

func main() {
	head := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 4,
					Next: &ListNode{
						Value: 5,
						Next: &ListNode{
							Value: 6,
							Next:  nil,
						},
					},
				},
			},
		},
	}
	fmt.Println(GetListNodeElements(head))
	reverse := reverseList(head)
	fmt.Println(GetListNodeElements(reverse))
}
