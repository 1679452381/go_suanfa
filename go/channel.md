## channel

* 给一个 nil channel 发送数据，造成永远阻塞

* 从一个 nil channel 接收数据，造成永远阻塞

* 给一个已经关闭的 channel 发送数据，引起 panic

* 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回 一个零值 和 false (ex: val,ok:= <-ch1)

* 无缓冲的channel是同步的，而有缓冲的channel是非同步的


**读写空阻塞，写关闭异常，都关闭零和假**