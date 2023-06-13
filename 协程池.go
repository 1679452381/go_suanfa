/**
 * @author X
 * @date 2023/6/13
 * @description
 **/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 协程池
type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// 初始化一个协程池
func NewPool(size int) *Pool {
	if size < 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add 增加一个执行
func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// Done完成一个执行 减1
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}
func (p *Pool) Wait() {
	p.wg.Wait()
}
func main() {
	pool := NewPool(5)
	fmt.Println("Goroutine begin nums:", runtime.NumGoroutine())
	for i := 0; i < 20; i++ {
		pool.Add(1)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("Goroutine continue nums:", runtime.NumGoroutine())
			pool.Done()
		}(i)
	}
	pool.Wait()
	fmt.Println("Goroutine done nums:", runtime.NumGoroutine())
}
