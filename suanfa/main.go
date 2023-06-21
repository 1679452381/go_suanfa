/**
 * @author X
 * @date 2023/6/19
 * @description
 **/
package main

import (
	"fmt"
	"sync"
)

//type App struct {
//	Messages map[string]string
//}

//发布消息

//订阅消息

func main() {
	var wg sync.WaitGroup
	msgChan := make(chan string, 10)
	//msgChan := make(chan string)
	wg.Add(2)
	//发布消息
	go func() {
		for i := 0; i < 10; i++ {
			msgChan <- fmt.Sprintf("消息%d", i)
		}
		close(msgChan)
	}()
	//消费者1
	go func() {
		for {
			select {
			case msg, ok := <-msgChan:
				if ok == false {
					fmt.Println("程序结束")
					wg.Done()
				}
				fmt.Println(msg)
				break
			default:
				break
			}
		}
	}()
	//消费者2
	go func() {
		for {
			select {
			case msg, ok := <-msgChan:
				if ok == false {
					fmt.Println("程序结束")
					wg.Done()
				}
				fmt.Println(msg)
				break
			default:
				break
			}
		}
	}()
	wg.Wait()
}
