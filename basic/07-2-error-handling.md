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

## 多Error处理
### 错误组合
### 链式执行
### 批量处理

## 错误感知

## 错误堆栈

## Web项目实践
