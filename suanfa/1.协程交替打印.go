package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	chanNum := make(chan int)
	chanLetter := make(chan int)
	index := 0
	i := 1
	a := "abcdefghijklmnopqrskuvwxyz"
	go func() {
		for {
			select {
			case <-chanNum:

				fmt.Println(i)
				i++
				chanLetter <- 1
				break
			default:
				break
			}
		}
	}()
	go func(wg *sync.WaitGroup) {
		for {
			select {
			case <-chanLetter:
				if index > len(a)-1 {
					wg.Done()
					return
				}
				fmt.Println(string(a[index]))
				index++
				chanNum <- 1
			default:
				break
			}
		}
	}(&wg)
	//chanNum <- 1
	chanLetter <- 1
	wg.Wait()

}
