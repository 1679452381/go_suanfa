## channel

```go
type hchan struct {
	qcount   uint           // 最多存储的元素数量
	dataqsiz uint           // 每个元素占 多大空间
	buf      unsafe.Pointer // 缓冲区 数组
	elemsize uint16 
	closed   uint32 //关闭状态
	elemtype *_type // 类型元数据
	sendx    uint   // 读下标位置
	recvx    uint   // 写下标位置
	recvq    waitq  // 写队列
	sendq    waitq  // 读队列
	
	lock mutex
}

```
* 给一个 nil channel 发送数据，造成永远阻塞

* 从一个 nil channel 接收数据，造成永远阻塞

* 给一个已经关闭的 channel 发送数据，引起 panic

* 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回 一个零值 和 false (ex: val,ok:= <-ch1)

* 无缓冲的channel是同步的，而有缓冲的channel是非同步的


**读写空阻塞，写关闭异常，都关闭零和假**


非阻塞式写法
```go
select {
case chan<-10:
	...
default:
	...
}
select {
case <-chan:
...
default:
...
}
```