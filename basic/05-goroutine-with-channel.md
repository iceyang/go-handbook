# 协程与通道（goroutine & channel）

通过Go协程，我们可以简单写出高并发的程序，而协程搭配通道，使得协程间的通讯变得简单。

而这也是`Go`语言的编程哲学：

```
Don’t communicate by sharing memory; share memory by communicating.
```

使用通信共享数据，而不要通过共享内存来通信。

下面通过一个简单的例子，我们来看看协程要如何方便的使用通道来通信。

## 生产者与消费者

这是一个非常简单的例子，我们用两个简单的函数，分别代表`生产者`和`消费者`，他们接受`channel`作为参数，生产者负责往里面增加数据，而消费者负责取出处理。

```Go
func producer(queue chan<- int) {
        for i := 0; i < 10; i++ {
                queue <- i
        }
        close(queue)
}

func consumer(queue <-chan int, finish chan int) {
        for v := range queue {
                fmt.Printf("Get value: %d \n", v)
                time.Sleep(time.Second)
        }
        finish <- 1
}

func main() {
        queue := make(chan int)
        finish := make(chan int)
        go producer(queue)
        go consumer(queue, finish)
        <-finish
}
```

在上面的代码中，函数`producer(queue chan<- int)`充当生产者，`consumer(queue <-chan int, finish chan int)`充当消费者。

在`main`函数中，将两个函数作为`goroutine`启动。同时加入通道变量`finish`，用于阻塞程序，让生产者和消费者有运行的时间。

> 在上面的例子，我们用了通道阻塞的办法，这是一种让主程序等待子协程的运行的办法。
