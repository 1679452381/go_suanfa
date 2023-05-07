package main

import "fmt"

type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{
		Queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyStack) Pop() int {
	val := this.Queue[len(this.Queue)-1]
	this.Queue = this.Queue[0 : len(this.Queue)-1]
	return val
}

func (this *MyStack) Top() int {
	return this.Queue[len(this.Queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}

func main() {
	stack := Constructor()
	fmt.Println(stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack)
	fmt.Println(stack.Top())
	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println(stack.Empty())
}
