# 协程与通道（goroutine & channel）

```
Don’t communicate by sharing memory; share memory by communicating.
```

这是`Go`语言的编程哲学，使用通信共享数据，而不要通过共享内存来通信。通过`Go`协程，我们可以简单写出高并发的程序，而协程搭配通道，使得协程间的通讯变得简单。

下面通过一个简单的`生产者与消费者`例子，我们来看看协程要如何方便的使用通道来通信。

## 生产者与消费者

用两个简单的函数，分别代表`生产者`和`消费者`，他们接受`channel`作为参数，生产者负责往里面增加数据，而消费者负责取出处理。

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

> 在上面的例子，我们用了通道finishl来阻塞主进程，这是一种让程序等待子协程运行的办法。
> 如果没有它，程序会很快执行完毕，goroutine没有得到运行的机会。

### 一个生产者与多个消费者

还是上面的例子，我们启动多个`consumer`，可以达到多个消费者进行处理的效果。

```Go
func consumer(queue <-chan int, finish chan int) {
        for v := range queue {
                fmt.Printf("Get value: %d \n", v)
                time.Sleep(time.Millisecond * 200)
        }
        finish <- 1
}

func main() {
        queue := make(chan int)
        finish := make(chan int)
        consumerCount := 3
        go producer(queue)
        for i := 0; i < consumerCount; i++ {
                go consumer(queue, finish)
        }
        for i := 0; i < consumerCount; i++ {
                <-finish
        }
}
```

多个`consume`的协程是并发执行的，相互之间互不干扰。
我们依旧使用`finish`通道来控制程序运行。

## 扩展

通过上面的例子，我们认识到通道在协程中可以起到传递信息的作用。

虽然例子中的通道只是简单的`chan int`类型，然而通道类型也可以是复杂的结构体，可以传递更加多的信息，可以延伸出许多更加实用的应用场景。

比如数据库io操作、比如作为web应用，又或者是一些需要并行计算的任务，在`Go`协程和通道的配合下，都有着许多可能性。

虽然`Go`开发起并发程序非常简单，但是也隐藏了许多复杂性，这是并发程序带来的问题，需要保证数据的原子性、一致性，需要考虑对数据的加锁，需要控制好协程间的配合或者执行顺序。这些相较于本文是比较复杂的内容，在后面的章节详细说明。
