package main

import "fmt"

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Init() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

func (q *MyQueue) Pop() int {
	inLen, outLen := len(q.stackIn), len(q.stackOut)
	//	判断stackOut 是否为空
	if outLen == 0 {
		//	都为空 返回-1
		if inLen == 0 {
			return -1
		}
		//	为空 将stackIn 全部输入到stackOut
		for i := inLen - 1; i >= 0; i-- {
			q.stackOut = append(q.stackOut, q.stackIn[i])
		}
		q.stackIn = []int{}
	}
	val := q.stackOut[len(q.stackOut)-1]
	q.stackOut = q.stackOut[0 : len(q.stackOut)-1]
	return val

}
func (q *MyQueue) Peek() int {
	val := q.Pop()
	if val == -1 {
		return -1
	}
	q.stackOut = append(q.stackOut, val)
	return val
}
func (q *MyQueue) Empty() bool {
	return len(q.stackOut) == 0 && len(q.stackIn) == 0
}

func main() {
	queue := Init()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue)
	fmt.Println(queue.Peek()) // 返回  1
	fmt.Println(queue)
	fmt.Println(queue.Pop()) // 返回  1
	fmt.Println(queue)
	fmt.Println(queue.Empty()) // 返回 false
}
