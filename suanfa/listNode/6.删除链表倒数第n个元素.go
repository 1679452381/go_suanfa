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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ret := &ListNode{
		Value: 0,
		Next:  head,
	}
	pre := ret
	cur := ret
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	for cur.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = pre.Next.Next
	return ret.Next

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
	head = removeNthFromEnd(head, 3)
	fmt.Println(GetListNodeElements(head))
}
