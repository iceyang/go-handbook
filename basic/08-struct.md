# 结构体

Go语言也拥有struct（结构体），结构体的使用，可以让我们将松散的数据字段组合起来，变成更有意义、可读性的类型。结构体可以拥有它自身对应的方法，像「面向对象编程」的「类」一样，拥有类的方法。

结构体也可以不拥有任何字段，只拥有方法，根据具体业务场景，它有它存在的意义。

## 定义
### 定义字段属性

要定义一个结构体，可以通过关键字`type`和`struct`完成：

```Go
type Person struct {
	name string
	age  int
}
```

我们定义了一个名字为`Person`的结构体。

### 定义方法

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

## 初始化结构体

下面是使用`Person`结构体的方式：

```Go
func TestStruct1(t *testing.T) {
	// 定义初始化，数值为默认值
	var person Person
	fmt.Println(person)

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

        // 使用 new() 函数初始化
        person4 := new(Person)
        person4.name = "Tim"
        fmt.Println(person4)
}
```

初始化时，我们可以通过`指定字段名`，或`按字段顺序`的方式初始化一个结构体的值。
要使用结构体的方法，可以通过`值.方法`，比如：`person.GetName()`。

> 当使用new函数初始化结构体时，得到的是一个指针变量，效果相等于 &Person{}。

## 方法的接收者

还记得上面我们为`Person`定义了两个`SetName`函数，它们的区别只在于接收者是不是指针。下面是一段简单示例以及运行结果，可以看出两个方法得到的效果不同：

```Go
/**
 * 执行结果，打印出来的是：
 * name:
 * SetName:  Justin
 * SetName2:  Justin
 */
func TestStruct2(t *testing.T) {
	person := Person{}
	fmt.Println("name:", person.GetName())
	person.SetName("Justin")
	fmt.Println("SetName: ", person.GetName())
	person.SetName2("Justin2")
	fmt.Println("SetName2: ", person.GetName())
}
```

结果出来了，我们初始化了一个变量`person`，一开始`name`属性为默认值`''`，因为我们没有赋予初始化的值。

在执行了第一个`SetName`之后，`person`的`name`变成了`Justin`，说明`name`被修改了。

而我们想通过`SetName2`去修改`name`的时候，却没有对`person`产生影响。

*原因*在于，当我们调用`SetName2`的时候，因为传递的对象是值对象，方法的调用会将`person`拷贝一份副本，我们对方法中`person`变量的修改，实际上不会对外面的调用者产生影响。而`SetName`赋予的是指针，对它的修改会直接操作到指针指向的对象值，所以它的修改是有效存在的。

## 属性的tag

在定义结构体属性字段的时候，可以为属性写上`tag`内容，比如：

```Go
type User struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	From   string             `bson:"from" json:"from"`
	Mobile string             `bson:"mobile" json:"mobile"`
}
```

跟在字段类型后面的""内容便是`tag`，`tag`最常用的场景就是为属性定义一些处理规则，比如进行`json`转换时，属性对应的`json`属性名可以自定义。

> `tag`的使用需要通过`reflect`包完成，而且它的定义必须符合`reflect`的规则，否则可能会导致读取的错误。
> 正常情况下，普通业务开发并不需要使用到reflect，主要还是由工具包、框架包进行实现。

## 匿名字段

匿名字段是指我们在定义结构体属性时没有为它指定名字，这时候字段名会与类型名一致，比如：

```Go
type AnonymousField struct {
        a int
        b int
        int
}

func TestStruct3(t *testing.T) {
	o := AnonymousField{}
	o.int = 10
	fmt.Println(o)
}
```

执行后我们可以看到，`o`的内容为：`{0 0 10}`，是因为AnonymousField拥有了字段`int`，同时我们对它进行了修改。

## 结构嵌套

还记得我们上面的例子中定义了`Person`结构体，我们现在定义`Adult`结构体内容如下：

```Go
type Person struct {
	name string
	age  int
}

type Adult struct {
	job string
	Person
}
```

`Adult`嵌套了一个匿名字段`Person`，在使用上，`Person`的字段可以被`Adult`直接使用，就好像使用`Adult`本身定义的字段一样：

```Go
func TestStruct4(t *testing.T) {
	adult := &Adult{}
	adult.job = "salesman"
	adult.name = "John"
	adult.age = 25
	fmt.Println(adult)
}
```

> 结构体的嵌套，使用上虽然类似面向对象的「继承」，但我们需要明确知道它们的区别。
>
> 嵌套 ≠ 继承，它代表的是一种组合的关系（has a）而非（is a）。

结构嵌套可以指定名称，比如：

```Go
type Person struct {
	name string
	age  int
}

type Adult struct {
	job    string
	person Person
}
```

不过在使用上就不能再直接使用`Person`的字段，需要`adult.person.name`，使用`adult.name`时会报错。

当结构嵌套的属性名或方法名产生了同名冲突，「内部」struct的字段会被「外部」struct覆盖，类似于「继承」的重写效果（override）。

## 总结

本文介绍了`struct`（结构体）的使用方式，包括：

* 定义
* 初始化
* 属性tag
* 匿名字段
* 结构嵌套

结构体的使用可以帮助我们将松散的数据整合在一起，让程序的可读性大大增强，编写业务代码更加方便。
