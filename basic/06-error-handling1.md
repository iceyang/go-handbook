# 错误处理（一）

在实际编程中，会出现各种各样的错误，对于预料之中的错误，我们可以优雅的处理，让程序平滑运行。而异常情况的错误，我们也可以预留保护机制，保证程序不会异常退出。

## 内置的接口类型：error

在 Go 语言里，我们可以利用提供的错误类型`error`来自定义错误。

error 是一个接口类型，是 Go 的内建类型，它的定义是：

```Go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
        Error() string
}
```

实现了`Error() string`方法的类型，都属于error接口类型。

当我们简单使用时，可以通过`errors.New()`来生成一个error，比如：

```Go
func division(divisor, dividend int) (int, error) {
        if dividend == 0 {
                return 0, errors.New("dividend cannot be zero")
        }
        return divisor / dividend, nil
}

func main() {
        if res, err := division(10, 5); err == nil {
                fmt.Println(res)
        }
        res, err := division(10, 0)
        if err == nil {
                fmt.Println(res)
        } else {
                fmt.Println(err)
        }
}
```

在上面的代码中，我们简单实现了一个整数除法，当被除数等于0时，会产生一个除数不能为0的error。

同时我们展示了一种函数处理错误返回值的方式：将错误声明作为返回的最后一个结果。
当执行完函数后，对err进行判断，假设err为`nil`时，则进行接下来的逻辑。
而当err不为`nil`，则说明执行发生了错误，这时需要对错误进行处理。

> 假如大家写过使用Promise之前的js，代码使用回调进行信息的传递，会有熟悉的感觉，通常会将err作为回调函数结果的第一位，作为约定的规范来处理。

### errors.New的实现

errors.New放在 Go errors 包中，下面是它的具体实现：

```Go
// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

实现很简单，返回了一个内置结构`errorString`的指针。而`errorString`结构只包含了一个字符串属性，记录了错误的信息，实现的`Error`方法是将该字符串返回。

## 如何对错误进行处理

通过前文我们知道，通常函数出现了错误，会将error作为返回值给到调用者，调用者接收到err不为nil时，需要做出相应的处理。通常的做法有：

1. 当我们确定error的类型值时，可以获取到它的潜在错误值，然后进行处理；
2. 当error都是由errors.New创建出来的，且已经有相应的初始化定义时，我们可以直接用判等的方式识别错误；
3. 当error的类型完全未知时，只能通过错误信息做判断。

上面的描述相对来说比较抽象，我们用具体的例子来做解释。

### 当确定error的类型值时

首先， 因为`error`是一个接口类型，只要实现了`Error() string`方法的实体，都可以被当成error，我们自定义两个错误类型：`ErrorOne`和`ErrorTwo`，它们的内容是：

```Go
type ErrorOne struct{}
type ErrorTwo struct{}

func (ErrorOne) Error() string {
        return "This is Error One"
}

func (ErrorTwo) Error() string {
        return "This is Error Two"
}
```

当我们调用的方法，可以确定它可能抛出什么类型的`error`时，可以用`switch`语句配合`.(type)`进行语句进行判断处理：

```Go
func handleErrorWithType(err error) {
        switch err.(type) {
        case *ErrorOne:
                fmt.Println("err is ErrorOne")
        case *ErrorTwo:
                fmt.Println("err is ErrorTwo")
        }
}

func doSomething1() error {
        num := rand.Intn(10)
        if num > 5 {
                return &ErrorOne{}
        } else if num < 5 {
                return &ErrorTwo{}
        }
        return nil
}

func handleErrorDemo1() {
        rand.Seed(time.Now().UnixNano())

        for i := 0; i < 10; i++ {
                if err := doSomething1(); err != nil {
                        handleErrorWithType(err)
                } else {
                        fmt.Println("There'is no error")
                }
        }
}
```

重点在于`handleErrorWithType`方法，由于我们知道产生的错误只能是`ErrorOne`和`ErrorTwo`两种类型，所以我们可以用`switch`进行判断，然后针对性进行处理。

> 这有点像其他语言的`try/catch`语句，当我们知道具体的类型时，可以直接catch具体的Error类型。

### 当error都是由errors.New创建出来的，且已经有相应的初始化定义时

```Go
var errorOne = errors.New("This is Error One")
var errorTwo = errors.New("This is Error Two")

func handleErrorWithType2(err error) {
	fmt.Print("handleErrorWithType2: ")
	switch err {
	case errorOne:
		fmt.Println("err is ErrorOne")
	case errorTwo:
		fmt.Println("err is ErrorTwo")
	}
}

func doSomething2() error {
	num := rand.Intn(10)
	if num > 5 {
		return errorOne
	} else if num < 5 {
		return errorTwo
	}
	return nil
}

func handleErrorDemo2() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		if err := doSomething2(); err != nil {
			handleErrorWithType2(err)
		} else {
			fmt.Println("There'is no error")
		}
	}
}
```

在`handleErrorWithType2`中，我们知道err的定义来源，而且我们拿得到引用，那么直接对它们做判断。

### 当error的类型完全未知时，只能通过错误信息做判断

假设我们不知道方法可能抛出的error接口究竟是什么内容，我们只能通过错误的内容进行判断。

```Go
func doSomething3() error {
	num := rand.Intn(10)
	if num > 5 {
		return &ErrorOne{}
	} else if num < 5 {
		return &ErrorTwo{}
	}
	return errors.New("Unknown Error")
}

func handleErrorWithType3(err error) {
	fmt.Print("handleErrorWithType3: ")
	errStr := err.Error()
	if strings.Contains(errStr, "Error One") {
		fmt.Println("err is ErrorOne")
		return
	}
	if strings.Contains(errStr, "Error Two") {
		fmt.Println("err is ErrorTwo")
		return
	}
	fmt.Println("unknown err")
}

func handleErrorDemo3() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		if err := doSomething3(); err != nil {
			handleErrorWithType3(err)
		} else {
			fmt.Println("There'is no error")
		}
	}
}
```

我们假设不知道`doSomething3`可能会抛出什么error，然后在`handleErrorWithType3`中，对错误信息进行比较。

虽然这种方法比较拙劣，但确实会有这样的情况发生。

## 总结

我们提到在`Go`语言中，也有内置的接口类型`error`，实现了`Error() string`方法的类型，都可以被当成`error`。

简单的`error`可以通过`Go errors`包的`New(text string)`方法生成。

在错误处理中，处理三种不同的情况：

1. 当我们确定error的类型值时，可以获取到它的潜在错误值，然后进行处理；
2. 当error都是由errors.New创建出来的，且已经有相应的初始化定义时，我们可以直接用判等的方式识别错误；
3. 当error的类型完全未知时，只能通过错误信息做判断。

在错误处理中，还包含着其他的问题，比如错误的上下文，错误信息如何传递，留待后续的文章说明。
