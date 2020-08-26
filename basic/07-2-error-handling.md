# 错误处理（二）

这一篇，主要是想讲一下在实践项目中处理错误的一些方式与思考。

接口类型`error`是 Go 的内置类型，实现了方法`Error() string`的类型，都可以被当成`error`。

现在，我们已经知道如何定义`error`，如何产生`error`，以及如何对`error`进行处理，参考 [错误处理（一）](./07-1-error-handling.md)

但是在真正的项目使用中，我们要如何组织错误，是 Gopher 一直探讨的话题。这里我会介绍我们在项目中的一些处理方式，以及如何组合这些错误处理方式，更优雅地组织代码。

## Error的返回方式
### 方法返回值

最常见的方式，便是将`error`作为返回值返回，也就是将如何处理交给方法调用者。这是许多 Golang 标准库的实现方式。
比如`os`包中的标准输出，调用`Write`方法，会返回`err`，在进行下一步操作时，当`err`如果不为`nil`时，调用者应该先对`err`进行处理。

```Go
_, err := os.Stdout.Write([]byte("Here is a string"))
if err != nil {
        fmt.Println(err)
}
```

再比如`github.com/go-redis/redis`库，在获取数据时，也需要对返回的错误进行识别：

```Go
result, err := cache.client.Get(key).Result()
if err != nil {
        // handle error
        return
}
fmt.Println(result)
```

总的来说，这是许多框架和标准库的实现方式，将错误的处理权上交给调用者。

回归到项目中，我们也可以在自己所写的方法中，将错误进行返回，一直到最上层的调用方，再将错误统一处理。

### panic抛出

通过 panic 抛出错误的方式，需要我们知道 Go panic/defer/recover 的机制，这里不累赘说明。
怎么实现呢？

1. 定义错误处理的defer方法
2. 在方法中使用recover()
3. 业务逻辑中panic出错误，就可以被recover捕捉

现在我们有两个方法，一个是业务逻辑DoSomething，一个是错误恢复Recover，将它们组合起来，可以起到处理错误的功能：

```Go
func Recover() {
        if err := recover(); err != nil {
                log.Printf("%s\n", err)
        }
}

func DoSomething() {
        defer Recover()
        log.Println("Here we go")
        panic(errors.New("An error occurs"))
        // unreachable code because of panic
        log.Println("Here we go2")
}
```

上面的方法我们是模拟了简单的业务处理。
真实项目中，panic的方式能让我们统一错误处理，比如作为Web项目的中间件使用。

## 多Error处理

上面提到了Error最常见的返回方式是通过返回值传递，当我们采用这种方式时，会遇到另外一种问题，也是经常被吐槽的点，
下面看一个具体的例子：

```Go
func foo() error {
        err := doSomething1()
        if err != nil {
                return err
        }
        err = doSomething2()
        if err != nil {
                return err
        }
        err = doSomething3()
        if err != nil {
                return err
        }
        return nil
}
```

典型的例子便是随处的`if err != nil`，`if`语句一用便是三行，让人难受。
对于这种问题，我做一些简单的探讨。

### 批量处理

还是用回上面的例子，假设`doSomething1`、`doSomething2`、`doSomething3`的结果相互之间不影响，也就是三个方法不是「事务」性操作，那么我们可以将返回的错误延迟到后面统一处理。

```Go
func dealErrors(errs ...error) error {
        // 处理错误
        return nil
}

func foo() error {
        err1 := doSomething1()
        err2 := doSomething2()
        err3 := doSomething3()

        err := dealErrors(err1, err2, err3)
        return err
}
```

因为3个方法之间互不影响，所以我们用方法`dealErrors`对它们返回的错误统一处理，`dealErrors`的具体处理方式由业务决定。

### 错误组合

看到上面`dealErrors`的例子，因为没有给出具体实现，所以还是处于一种比较抽象的状态。
对于这种情况，有一种很常见的处理方式，便是把错误组合起来返回。

现在我们定义一个包，叫`multierr`，并提供方法：

* func Combine(errors ...error) error
* func Errors(err error) []error

前者将错误组合起来变成一个新的error，后者可以将前者产生的error进行拆包。我们在这里不关注实现的细节，只关注它的行为。

那么，现在针对上面的例子，我们可以这样写：

```Go
func foo() error {
        return multierr.Combine(
          doSomething1(),
          doSomething2(),
          doSomething3(),
        )
}

func main() {
        err := foo()
        errors := multierr.Errors(err)
        if len(errors) > 0 {
                fmt.Println("The following errors occurred:")
                for _, e := range errors {
                        fmt.Println(e)
                }
        }
}
```

在 foo 中，执行了三个逻辑方法，它们都可能产生错误，然后我们使用 multierr.Combine 将错误组合起来后返回。

在 main 中，捕获 foo 返回的错误，并使用 multierr.Errors 进行拆包，假设有错误产生，便将他们输出打印。

对于这个multierr包，实际上是 uber 的一个实现，有兴趣的可以查看 -> [multierr](https://github.com/uber-go/multierr)

### 链式执行

还是讨论上面的例子，我们使用了`mutierr.Combine(errors ...error)`，对多个错误进行了合并。`foo`中的三个逻辑方法，在调用Combine的时候，已经执行完毕，Combine只不过是对三个error进行处理。

能这么做，是因为`foo`中的三个逻辑方法互不影响，假设三个方法是有顺序依赖的，或者说，一个方法执行失败，那后续的方法也没必要执行，这个时候我们或许可以这么做：

我们定义一个多方法处理的函数Execute。

```Go
type Handler func() error

func Execute(handlers ...Handler) error {
	for _, handler := range handlers {
		if err := handler(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := Execute(
		func() error {
			fmt.Println("Here we are1")
			return nil
		},
		func() error {
			fmt.Println("Here we are2")
			return errors.New("error occurs")
		},
		func() error {
			fmt.Println("Here we are3")
			return nil
		},
	)
	fmt.Println(err)
}
```

上面的代码，简单的定义了一个`Handler`用于代表会返回一个`error`参数的方法，`Execute`则负责将批量的`Handler`进行执行。
只要其中一个`Handler`产生了错误，那么执行链便会终止。

运行上述代码，会看到，`Here we are3`没有输出，因为它并不会被执行。

> 在这里，其实还有很多讨论的空间。比如说，那链式调用，上下文怎么办，后边的方法经常会依赖到前边方法的返回值。
>
> Go 对于错误作为返回值的处理，其实跟JavaScript的回调处理十分相似，也是约定俗成的错误返回。
>
> 因为这里只是展示错误处理的一种可能性，所以如果对这种方式有兴趣，不妨看看js非常有名的 [async](https://github.com/caolan/async) 包，可以说是非常经典的一个工具了。

## 错误感知

错误感知只是我对于下面想介绍的这种错误处理方式的一个直观叫法，也是探讨之一。

在许多语言、许多框架中，经常会见到一种设计：**链式调用**，它的一种经典实现是`建造者模式`，比如：

```Go
type User struct {
        Name string
        Age  int
        Sex  int
}

// 假设有一个Builder，可以帮助我们构造User
u := Builder().Name("Justin").Age(11).Sex(1)
// 我们使用了Builder来构造一个User对象。
```

那回归到我们的实际业务逻辑中，我们想要把链式调用用到其他地方，并且对错误进行比较好的处理，要怎么做呢？

打个比方，现在有一个业务逻辑方法，叫`foo`，`foo`由三个粒度更小的方法`bar1`、`bar2`、`bar3`组成，三个方法直接又有先后顺序，当采用链式调用时，会长成这样子：

```Go
func foo() {
        err := bar1().bar2().bar3()
        if err != nil {
                fmt.Println(err)
        }
}
```

我们想达到的目的是，当`bar1`产生错误时，我们可以先不进行处理，直接调用`bar2`。
在`bar2`中，因为它们处于同一个上下文，在设计上`bar2`是可以知道`bar1`是否产生了错误，那么会有两种结果：

* bar1 没有错误产生，那么 bar2 照常运行
* bar1 产生了错误，bar2 不执行逻辑，将错误抛出。

在`bar3`上，也采用相同的逻辑思路进行处理。

这种方式，暂且叫错误感知吧。
它方便了写法，但也对业务上下文要求高许多，同时会有比较大的耦合，使用较为局限，在我看来，它比较适合数据库、缓存、请求库之类的框架使用。

## 总结

以上是对`Go`错误处理的一些更广范围的思考与运用，希望能有所帮助。
