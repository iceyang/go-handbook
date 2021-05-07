# errgroup - 捕获 goroutine 的错误

在`Go`里面，想要让程序并发运行是一件很简单的事情，只要在func执行的时候，加上`go`关键字便实现了。 
但同时又引入了另外一个问题，`go`执行我们捕获不到错误信息。

对于普通错误信息，我们可能只需要将其打到 log 日志，然后到日志收集系统去查询。

另外还可以想到用`go`的搭档`channel`，通过通道将错误传递出来。

本文主要介绍另外一种方式，通过Go官方提供的扩展包errgroup传递，它是基于 WaitGroup 实现的。

## 如何使用errgroup
`errgroup`位于`golang.org/x/sync/errgroup`，使用起来非常简单，因为它只提供了两个方法便达到了目的，这里有一个使用例子：

```go
func doSomethings(index int) error {
	fmt.Printf("Job[%d] finished\n", index)
	if rand.Intn(5) >= 4 {
		return fmt.Errorf("an error occurs: %d", index)
	}
	return nil
}

// 即使出错了也会跑完所有任务
func TestErrGroup1(t *testing.T) {
	group, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 10; i++ {
		index := i
		group.Go(func() error {
			time.Sleep(time.Duration(index) * time.Second)
			return doSomethings(index)
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}
```

上面的代码中，主要经历下面几个步骤：
1. 程序通过`errgroup.WithContext(context.Background())`创建了一个`group`
2. 然后通过`group.Go`启动协程，总共启动了10个任务。在方法`doSomethings`中，有可能返回error。
3. 最后通过`group.Wait()`阻塞整个程序，直到所有任务执行完。

下面是一次运行结果：

```shell
 ~/project/github/go-handbook/code/sync/ [master*] go test --run TestErrGroup1
Job[0] finished
Job[1] finished
Job[2] finished
Job[3] finished
Job[4] finished
Job[5] finished
Job[6] finished
Job[7] finished
Job[8] finished
Job[9] finished
an error occurs: 4
```

可以看到，任务[4]发生了错误，而程序依旧执行了10个job。

> 这里需要注意的是：
> group.Wait()会返回第一个发生的错误，所以这里不代表5-9的Job是没有错误发生的。
> 同时group在得知发生错误之后，会通过`context`告知cancel状态，所以如果我们想要中断其他任务，则需要在创建Group之时，将`context`一并获取。

## 发生错误时，停止其他goroutine
在上面的例子中，某一个任务发生错误后，其实Group已经通过`context`告知了cancel状态，但由于我们没有接收创建Group时的`context`，导致其他任务依旧会执行（当然，也可能任务之间并没有关联，出错了不影响其他任务的执行，这时候可以忽略）。当我们需要取消其他任务时，只需要通过`context`，就可以感知到，比如：

```go
// 可以感知到错误而停止其他任务
func TestErrGroupWithCancel(t *testing.T) {
	group, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 5; i++ {
		index := i
		rand.Seed(time.Now().UnixNano())
		sleepTime := rand.Intn(5)
		group.Go(func() error {
			select {
			case <-time.After(time.Duration(sleepTime) * time.Second):
				return doSomethings(index)
			case <-ctx.Done():
				return ctx.Err()
			}
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}
```

这里跟上面第一个例子的区别在于，我们通过`select`关键字，监听创建`group`时返回的`ctx`，通过`ctx.Done()`，我们可以知道其他任务是否发生了错误。

## 给errgroup加上超时context
既然我们可以通过`ctx`得知`cancel`状态，那如果我们需要给`errgroup`增加时间限制，那只需要创建`WithTimeout`的`context`就可以做到了：

```go
// 超时中断
func TestErrGroupWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		index := i
		group.Go(func() error {
			select {
			case <-time.After(time.Duration(index) * time.Second):
				// 模拟耗时操作
				fmt.Printf("finished:%d\n", index)
				return nil
			case <-ctx.Done():
				fmt.Printf("canceled:%d\n", index)
				return ctx.Err()
			}
		})
	}
	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}
```

## 总结

我们介绍了`errgroup`如何捕获到`goroutine`的错误。
同时介绍通过上下文，可以让`goroutine`获取`cancel`状态，另外如果需要超时控制，则在`context`创建之时，给与`WithTimeout`即可。

总体来讲，官方包提供的`errgoup`还是比较轻量级，对于上下文的处理，可能还是需要使用者花多点功夫。
另外`Wait()`方法返回的只是第一个错误，有时可能依旧满足不了需求，需要加于扩展。

本文的具体代码可以在 [此处](../code/sync) 获得。
