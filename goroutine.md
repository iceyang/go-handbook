# 协程 goroutine

我们前面介绍过`channel`数据类型，它是传递消息的载体，是并发安全的。Go 语言采用了 CSP 的并发模型，除了`channel`外，`goroutine`也是它的基础。那`goroutine`是什么？

要解答这个问题，我们需要先了解下进程与线程。

> `goroutine`是`Go`语言并发模型的基础，是并发执行的实体。
