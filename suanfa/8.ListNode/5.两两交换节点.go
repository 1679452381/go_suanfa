package main

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
//func swapPairs(head *ListNode) *ListNode {
//	ret := &ListNode{
//		Value: 0,
//		Next:  head,
//	}
//	cur := ret
//	for cur.Next != nil && cur.Next.Next != nil {
//		temp := cur.Next.Next
//		pre := cur.Next
//
//		pre.Next = temp.Next
//		temp.Next = pre
//		cur.Next = temp
//
//		cur = pre
//	}
//	return ret.Next
//}
//
//func main() {
//	head := &ListNode{
//		Value: 1,
//		Next: &ListNode{
//			Value: 2,
//			Next: &ListNode{
//				Value: 3,
//				Next: &ListNode{
//					Value: 4,
//					Next: &ListNode{
//						Value: 5,
//						Next: &ListNode{
//							Value: 6,
//							Next:  nil,
//						},
//					},
//				},
//			},
//		},
//	}
//	fmt.Println(GetListNodeElements(head))
//	head = swapPairs(head)
//	fmt.Println(GetListNodeElements(head))
//}
