package main

//
//import "fmt"
//
//type ListNode struct {
//	Value int
//	Next  *ListNode
//}
//
//// 遍历链表
//func GetListNodeElements(head *ListNode) []int {
//	cur := head
//	listNodeValues := make([]int, 0)
//	for cur != nil {
//		listNodeValues = append(listNodeValues, cur.Value)
//		cur = cur.Next
//	}
//	return listNodeValues
//}
//

//// 输入一个slice创建一个链表
//func SetListNodeBySlice(arr []int) *ListNode {
//	ret := &ListNode{
//		Value: 0,
//		Next:  nil,
//	}
//	cur := ret
//	for _, val := range arr {
//		cur.Next = &ListNode{
//			Value: val,
//			Next:  nil,
//		}
//		cur = cur.Next
//	}
//	return ret.Next
//}

//// 链表长度
//func GetListNodeLen(head *ListNode) int {
//	cur := head
//	count := 0
//	for cur != nil {
//		count++
//		cur = cur.Next
//	}
//	return count
//}
//
//// 获取链表第index个节点的值
//func (head *ListNode) GetNodeByIndex(index int) int {
//	if index == 0 {
//		return -1
//	}
//	cur := head
//	count := 0
//	for cur != nil {
//		count++
//		if index == count {
//			return cur.Value
//		}
//		cur = cur.Next
//	}
//	return -1
//}
//
//// 在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点
//func (head *ListNode) AddAtHead(val int) *ListNode {
//	ret := &ListNode{
//		Value: 0,
//		Next:  head,
//	}
//	ret.Next = &ListNode{
//		Value: val,
//		Next:  nil,
//	}
//	ret.Next.Next = head
//	return ret.Next
//}
//
//// 将值为 val 的节点追加到链表的最后一个元素
//func (head *ListNode) AddAtTail(val int) *ListNode {
//	ret := &ListNode{
//		Value: 0,
//		Next:  head,
//	}
//	newNode := &ListNode{
//		Value: val,
//		Next:  nil,
//	}
//	cur := ret
//	for cur.Next != nil {
//		cur = cur.Next
//	}
//	cur.Next = newNode
//
//	return ret.Next
//}
//
//// 在链表中的第 index 个节点之前添加值为 val  的节点。
//// 如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//func (head *ListNode) AddAtIndex(index, val int) *ListNode {
//	if index < 0 {
//		index = 1
//	}
//	newNode := &ListNode{val, nil}
//	ret := &ListNode{
//		Value: 0,
//		Next:  head,
//	}
//	cur := ret
//	for i := 0; i < index-1; i++ {
//		cur = cur.Next
//	}
//	temp := cur.Next
//	cur.Next = newNode
//	newNode.Next = temp
//	return ret.Next
//}
//
//// 如果索引 index 有效，则删除链表中的第 index 个节点。
//func (head *ListNode) DeleteByIndex(index int) bool {
//	if index == 0 {
//		return false
//	}
//	ret := &ListNode{
//		Value: 0,
//		Next:  head,
//	}
//	count := 0
//	for ret.Next != nil {
//		count++
//		if index == count {
//			ret.Next = ret.Next.Next
//		}
//		ret = ret.Next
//	}
//	return false
//}
//
//func main() {
//	head := &ListNode{
//		Value: 1,
//		Next: &ListNode{
//			Value: 2,
//			Next: &ListNode{
//				Value: 6,
//				Next: &ListNode{
//					Value: 3,
//					Next: &ListNode{
//						Value: 4,
//						Next: &ListNode{
//							Value: 5,
//							Next: &ListNode{
//								Value: 6,
//								Next:  nil,
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//	fmt.Println(head, GetListNodeLen(head), GetListNodeElements(head))
//	fmt.Println(head.GetNodeByIndex(0))
//
//	//在头部添加节点 99
//	head = head.AddAtHead(99)
//	fmt.Println(head, GetListNodeLen(head), GetListNodeElements(head))
//
//	//在尾部添加99
//	head = head.AddAtTail(99)
//	fmt.Println(head, GetListNodeLen(head), GetListNodeElements(head))
//
//	//在第3位添加99
//	head = head.AddAtIndex(3, 99)
//	fmt.Println(head, GetListNodeLen(head), GetListNodeElements(head))
//
//	//删除第5个元素
//	head.DeleteByIndex(5)
//	fmt.Println(head, GetListNodeLen(head), GetListNodeElements(head))
//}
