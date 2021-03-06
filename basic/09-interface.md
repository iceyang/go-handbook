# 接口（interface）

Go 语言没有类与继承的概念，但它提供的「接口」功能，却能让我们优雅地实现面向对象的特性。

在许多面向对象语言中，接口定义完之后，要实现一个接口，需要显示写明，比如`Java`使用`implements`关键字来实现接口。
而 Go 语言中的接口是隐式的，非侵入性实现，方便又灵活。

它允许我们提供新的接口类型，却无需改动到旧的实现代码。
接口的定义者，只需要定义好接口内容，不需要知道会被哪些类型实现。而接口的实现者，也只需要知道实现了哪个接口，不用显示指明。

## 声明（定义）接口

通过关键字`type`和`interface`，可以定义*接口*，*接口*可以包含一个或数个方法声明。
对于方法声明，跟普通方法一样，只是它不需要提供实现，另外方法的参数名以及返回值的名称可以写也可以不写。

比如我们常用的`fmt`包中，就有这么一个接口：

```Go
type Stringer interface {
	String() string
}
```

`Stringer`是接口的名称，它拥有一个`String() string`的方法。

## 如何实现接口

要对接口进行实现，需要满足两个条件

1. 实现的接口方法，签名必须一致，也就是方法的「名称」、「参数列表」，「返回值列表」都得一致；
2. 实现了接口的所有方法。

只要同时符合上述两个条件，类型便实现了接口。

还是以`Stringer`接口为例，只要类型实现了`String() string`方法，那么它就可以被当成`Stringer`使用。

在`fmt`的输出中，有这么一段代码：

```Go
// If a string is acceptable according to the format, see if
// the value satisfies one of the string-valued interfaces.
// Println etc. set verb to %v, which is "stringable".
switch verb {
case 'v', 's', 'x', 'X', 'q':
        // Is it an error or Stringer?
        // The duplication in the bodies is necessary:
        // setting handled and deferring catchPanic
        // must happen before calling the method.
        switch v := p.arg.(type) {
        case error:
                handled = true
                defer p.catchPanic(p.arg, verb, "Error")
                p.fmtString(v.Error(), verb)
                return

        case Stringer:
                handled = true
                defer p.catchPanic(p.arg, verb, "String")
                p.fmtString(v.String(), verb)
                return
        }
}
```

> 上面代码涉及到接口的断言，后续会说明

在代码中，`fmt`包通过类型判断输出的对象是否属于`Stringer`，如果是的话，会调用它的`String()`方法。
我们现在提供`Person`类型，实现了`String() string`，看看它的效果：

```Go
type Person struct {
	name string
}

func (person *Person) String() string {
	return "Person: " + person.name
}

func TestFmt(t *testing.T) {
	person := &Person{"Justin"}
	fmt.Println(person)
}
```

运行后能看到输出为：
```
Person: Justin
```

说明`Person`是被当成`Stringer`使用的。

## 类型与接口的关系

因为 Go 中的接口与实现是隐式关系，这带来了极大的灵活性：

* 一个类型可以同时实现多个接口，而且接口之间可以完全没有关系，相互独立；
* 在不改动原有类型的代码的基础上，我们可以提取出新的接口。

比如现在有一组类型，它们都拥有`func Fly()`，那我们可以提取出：

```Go
type Flyer interface {
        Fly()
}
```

那原有的拥有Fly方法的类型，都可以被当成Flyer使用。

### 类型断言

有时候我们想把接口转换成具体的类型或者其他接口，可以使用类型断言：

```Go
value, ok := x.(T)
```

效果是将`x`转换成具体类型`T`。
假设转换成功，那么`value`将会被赋值，类型是`T`，同时`ok`值为true；如果失败了，那`value`不会被赋值，`ok`值为false。

我们可以不接收`ok`的值，比如：

```Go
value := x.(T)
```

转换效果与接收`ok`是一致的，不同点在于，假设转换失败，那么程序会直接`panic`。

`switch`还可以配合`.(type)`一起使用，针对具体的类型，做出不同的逻辑处理。

```Go
func assertType(i interface{}) {
        switch i.(type) {
        case string:
                fmt.Println("Type is string")
        case error:
                fmt.Println("Type is error")
        default:
                fmt.Println("Unknown type")
        }
}
```

## 接口嵌套

类型可以嵌套，接口也可以嵌套。
比如Go语言的io包中，有这么几个接口：

```Go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
```

它们可以组合成各种各样的接口：

```Go
// ReadWriter is the interface that groups the basic Read and Write methods.
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser is the interface that groups the basic Read and Close methods.
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser is the interface that groups the basic Write and Close methods.
type WriteCloser interface {
	Writer
	Closer
}

// ReadWriteCloser is the interface that groups the basic Read, Write and Close methods.
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// ...
```

当一个类型既实现了`Reader`，也实现了`Writer`，那么它便是`ReadWriter`，以此类推。

有了这个特性，我们在定义跟使用接口时也更为灵活，每一项特性都可以是一个小的`interface`，多个接口可以组合成新的接口。

## 空接口

空接口是指没有任何方法的接口，它是接口的特殊形式，所有类型都属于空接口。

它类似于 Java 中的 Object，C语言的 void*，TypeScript 中的 any。有需要它的场景，但不能滥用。

首先我们看看空接口可以怎么接收参数。

```Go
type Person struct {
	name string
}

func show(entry interface{}) {
	fmt.Println(entry)
}

func TestVoidInterface(t *testing.T) {
	show(1)
	show("Hello World")
	show(Person{"Justin"})
}

// 执行结果
// 
// === RUN   TestVoidInterface
// 1
// Hello World
// {Justin}
// --- PASS: TestVoidInterface (0.00s)
```

我们定义了一个`show`函数，入参为`entry interface{}`，它可以接收任何参数并将其打印出来。

那我们要怎么从空接口来获取具体类型？
答案是 **使用断言**。

比如：

```Go
var a interface{} = 100
var b int = a.(int)
fmt.Println(b)
_, ok := a.(string)
if !ok {
        fmt.Println("a is not a string")
}
```

记得断言可以有两个接收参数，第二个参数是判断是否转换成功用的。
