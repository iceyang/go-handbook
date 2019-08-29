# 协程与通道（goroutine & channel）

通过Go协程，我们可以简单写出高并发的程序，而协程搭配通道，使得协程间的通讯变得简单。

而这也是`Go`语言的编程哲学：

```
Don’t communicate by sharing memory; share memory by communicating.
```

通过通信共享数据，不要通过共享内存来通信。

通过一个「生产者与消费者」的例子，我们来看看协程要如何方便的使用通道来通信。

## 生产者与消费者
