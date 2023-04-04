package main

//
//import (
//	"container/heap"
//	"time"
//)
//
//var cache = map[int]*Node
//
//const LIFE = 10
//
//
//func init() {
//	cache = make(map[int]*Node)
//}
//
//type Node struct {
//	DeadLine int64
//	Key int
//}
//
//type TimeoutHeap []*Node
//
//func (pq TimeoutHeap)Len() int {
//	return  len(pq)
//}
//
//func (pq TimeoutHeap)Less(i,j int) bool {
//	return pq[i].DeadLine < pq[j].DeadLine
//}
//
//func (pq TimeoutHeap)Swap(i,j int)  {
//	pq[i],pq[j] = pq[j],pq[i]
//}
//
//func (pq *TimeoutHeap)push(x interface{})  {
//	item := x.(*Node)
//	*pq = append(*pq,item)
//}
//
//func (pq *TimeoutHeap)pop() interface{} {
//	old := *pq
//	len := pq.Len()
//	res := old[len-1]
//	*pq = old[0:len-1]
//	return res
//}
//
//func testTimeoutCache() {
//	pq:=make(TimeoutHeap,0,5)
//	heap.Init(&pq)   //从无序状态构建堆
//
//	for i := 0; i < 10; i++ {
//		node := &Node{DeadLine: time.Now().UnixNano()+LIFE,Key: i}
//		cache[i] = nodes
//		heap.Push(&pq,node)
//		time.Sleep(20*time.Millisecond)
//	}
//
//}
