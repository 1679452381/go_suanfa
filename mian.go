/**
 * @author X
 * @date 2023/6/13
 * @description
 **/
package main

import "fmt"

var ch chan int

func main() {
	//s := []int{1, 2, 3, 4, 5}
	//for _, v := range s {
	//	s = append(s, v)
	//	fmt.Printf("len(s)=%v\n", len(s))
	//}

	//var s1 []int
	//s2 := make([]int, 0)
	//s4 := make([]int, 0)
	//
	//fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	//fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	//fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)

	//wg := sync.WaitGroup{}
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		fmt.Println(i)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	//fmt.Println("exit")

	//	限制goroutine数量
	ch = make(chan int, 5)
	for i := 0; i < 10; i++ {
		ch <- 1
		fmt.Println("ch value send", ch)
		go elegance()
		fmt.Println("result i:", i)
	}
}

func elegance() {
	<-ch
	fmt.Println("ch value receive", ch)
}
