# 结构体与接口

## 结构体

结构体的使用，可以让我们将松散的数据字段组合起来，变成更有意义、可读性的类型。结构体可以拥有它自身对应的方法，像「面向对象编程」的「类」一样，拥有类的方法。

结构体也可以不拥有任何字段，只拥有方法。

### 定义

要定义一个结构体，可以通过关键字`type`和`struct`完成：

```Go
type Person struct {
	name string
	age  int
}

func (p *Person) GetName() string {
	return p.name
}
```

我们定义了一个名字为`Person`的结构体，它拥有方法`GetName()`。

> 定义结构体的方法，与函数定义类似，不同的是在`func`和`方法名`中间，需要指定函数的接收者是谁，通过上面的定义后，我们就可以将函数作为结构体的方法来使用。

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
