# 错误处理

在实际编程中，会出现各种各样的错误，对于预料之中的错误，我们可以优雅的处理，让程序平滑运行。而异常情况的错误，我们也可以预留保护机制，保证程序不会异常退出。

## error类型

在 Go 语言里，我们也可以利用提供的错误类型error来对错误进行处理。

error 是一个接口类型，是 Go 的内建类型，我们可以通过`errors.New()`来生成一个error，比如：

```Go
func division(divisor, dividend int) (int, error) {
        if dividend == 0 {
                return 0, errors.New("dividend cannot be zero")
        }
        return divisor / dividend, nil
}

func main() {
	if res, err := division(10, 5); err != nil {
		fmt.Println(res)
	}
	if res, err := division(10, 0); err != nil {
		fmt.Println(res)
	}
}
```

在上面的代码中，我们简单实现了一个整数除法，当被除数等于0时，会产生一个除数不能为0的error。

同时我们展示了一种函数处理错误返回值的方式：将错误声明作为返回的最后一个结果。
当执行完函数后，对err进行判断，假设err为`nil`时，则进行接下来的逻辑。
而当err不为`nil`，则说明执行发生了错误，这时需要对错误进行处理。

> 假如大家写过使用Promise之前的js，代码使用回调进行信息的传递，会有熟悉的感觉，通常会将err作为回调函数结果的第一位，作为约定的规范来处理。

## 如何对错误进行处理

TODO
