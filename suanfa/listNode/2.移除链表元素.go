/**
 * @author X
 * @date 2023/5/4
 * @description
 **/
package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

// 遍历链表
func GetListNode(head *ListNode) []int {
	cur := head
	listNodeValues := make([]int, 0)
	for cur != nil {
		listNodeValues = append(listNodeValues, cur.Value)
		cur = cur.Next
	}
	return listNodeValues
}

// 链表长度
func GetLen(head *ListNode) int {
	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

// 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	ret := &ListNode{
		0,
		head,
	}
	cur := ret
	for cur.Next != nil {
		if cur.Next.Value == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return ret.Next
}

func main() {
	head := &ListNode{
		Value: 1,
		Next: &ListNode{
			Value: 2,
			Next: &ListNode{
				Value: 6,
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
		},
	}
	fmt.Println(head, GetLen(head), GetListNode(head))
	removeElements(head, 6)
	fmt.Println(head, GetLen(head), GetListNode(head))

}
