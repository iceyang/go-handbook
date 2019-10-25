# 结构体与接口

## 结构体

结构体的使用，可以让我们将松散的数据字段组合起来，变成更有意义、可读性的类型。结构体可以拥有它自身对应的方法，像「面向对象编程」的「类」一样，拥有类的方法。

结构体也可以不拥有任何字段，只拥有方法，根据具体业务场景，它有它存在的意义。

### 定义结构体

要定义一个结构体，可以通过关键字`type`和`struct`完成：

```Go
type Person struct {
	name string
	age  int
}
```

我们定义了一个名字为`Person`的结构体。

### 定义结构体的方法

```Go
func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) SetName2(name string) {
	p.name = name
}
```

定义结构体的方法，与函数定义类似，不同的是在`func`和`方法名`中间，需要指定函数的接收者是谁，通过上面的定义后，我们就可以将函数作为结构体的方法来使用。

我们为`Person`定义了3个方法：`GetName() string`，`SetName(name string)`, `SetName2(name string)`。

> 观察两个SetName方法，它们的区别在于接收者是谁，一个是Person类型，一个是Person类型的指针。
> 后面会说明它们会有什么不同。

### 初始化结构体

下面是使用`Person`结构体的方式：

```Go
func TestStruct1(t *testing.T) {
	// 通过字段名称初始化
	person1 := Person{
		name: "Justin",
		age:  18,
	}
	fmt.Println(person1.GetName())

	// 按字段顺序初始化
	person2 := Person{"Justin", 18}
	fmt.Println(person2)

	// 部分初始化
	person3 := Person{
		name: "Justin",
	}
	fmt.Println(person3)
}
```

初始化时，我们可以通过`指定字段名`，或`按字段顺序`的方式初始化一个结构体的值。
要使用结构体的方法，可以通过`值.方法`，比如：`person.GetName()`。

### 方法的接收者

还记得上面我们为`Person`定义了两个`SetName`函数，它们的区别只在于接收者是不是指针。下面是一段简单示例以及运行结果，可以看出两个方法得到的效果不同：

```Go
func TestStruct2(t *testing.T) {
	person := Person{}
	fmt.Println("name:", person.GetName())
	person.SetName("Justin")
	fmt.Println("SetName: ", person.GetName())
	person.SetName2("Justin2")
	fmt.Println("SetName2: ", person.GetName())
}

/**
 * 执行结果，打印出来的是：
 * name:
 * SetName:  Justin
 * SetName2:  Justin
 */
```

结果出来了，我们初始化了一个变量`person`，一开始`name`属性为默认值`''`，因为我们没有赋予初始化的值。

在执行了第一个`SetName`之后，`person`的`name`变成了`Justin`，说明`name`被修改了。

而我们想通过`SetName2`去修改`name`的时候，却没有对`person`产生影响。

*原因*在于，当我们调用`SetName2`的时候，因为传递的对象是值对象，方法的调用会将`person`拷贝一份副本，我们对方法中`person`变量的修改，实际上不会对外面的调用者产生影响。而`SetName`赋予的是指针，对它的修改会直接操作到指针指向的对象值，所以它的修改是有效存在的。
