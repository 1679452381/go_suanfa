### mutex 互斥锁

```go
package sync

type Mutex struct {
	state int32   // 互斥锁状态 默认值为0 为加锁
	sema  uint32  //信号量 等待队列
}
```

* mutex饥饿模式
  * 正常模式-> 饥饿模式
    * 当一个goroutine本次加锁等待时间超过1ms 
  * 饥饿模式 -> 正常模式
    * goroutine等待时间小于1ms
    * goroutine是最后一个等待者
* Lock
* unLock