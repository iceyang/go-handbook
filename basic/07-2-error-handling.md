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

因为3个方法之间互不影响，所以我们用方法`dealErrors`对它们返回的错误统一处理，具体处理方式由业务决定。

### 错误组合

### 链式执行

## 错误感知

## 错误堆栈

## Web项目实践
