package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

// 遍历链表
func GetListNodeElements(head *ListNode) []int {
	listNodeValues := make([]int, 0)
	if head == nil {
		return listNodeValues
	}
	cur := head
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

// 输入一个slice创建一个链表
func SetListNodeBySlice(arr []int) *ListNode {
	ret := &ListNode{
		Value: 0,
		Next:  nil,
	}
	cur := ret
	for _, val := range arr {
		cur.Next = &ListNode{
			Value: val,
			Next:  nil,
		}
		cur = cur.Next
	}
	return ret.Next
}

// 找两个链表的相交处
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headALen := GetLen(headA)
	headBLen := GetLen(headB)
	diff := 0
	curA := headA
	curB := headB
	var fast, slow *ListNode
	if headALen > headBLen {
		diff = headALen - headBLen
		fast = curA
		slow = curB
	} else {
		diff = headBLen - headALen
		fast = curB
		slow = curA
	}
	for i := 0; i < diff; i++ {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func main() {
	headA := SetListNodeBySlice([]int{4, 1, 8, 4, 5})
	headB := SetListNodeBySlice([]int{5, 0, 1, 8, 4, 5})
	fmt.Println(GetListNodeElements(headA), GetListNodeElements(headB))
	node := getIntersectionNode(headA, headB)
	if node != nil {
		fmt.Printf("ntersected at '%c'", node.Value)
	} else {
		fmt.Println(node)
	}
	fmt.Println(GetListNodeElements(headA), GetListNodeElements(headB))

}
