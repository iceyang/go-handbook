# Go语言的多态

在之前的文章里，我们了解过 [struct](/golang/struct/) 与 [interface](/golang/interface/)。

通过结构体，我们可以将松散的数据字段组合成有意义的类型，同时它可以拥有自己的方法，让我们可以像面向对象的「类」一样去使用。

而通过接口，我们可以统一 **类型与行为**，非侵入性的实现允许在完全不需要改动旧代码的情况下，让 struct **实现** interface。

还记得面向对象中有另外一个特性：多态。

多态在 Go 语言中，是如何体现呢，我们先看一段简单的例子：

```Go
type Animal interface {
        Eat()
        Sleep()
}

type Cat struct {
        Name string
}

func (c Cat) Eat() {
        fmt.Printf("Cat %s is eating.\n", c.Name)
}

func (c Cat) Sleep() {
        fmt.Printf("Cat %s is sleeping.\n", c.Name)
}

type Dog struct {
        Name string
}

func (d Dog) Eat() {
        fmt.Printf("Dog %s is eating.\n", d.Name)
}

func (d Dog) Sleep() {
        fmt.Printf("Dog %s is sleep.\n", d.Name)
}

func AnimalEat(animal Animal) {
        animal.Eat()
}

func TestPolymorphismOne(t *testing.T) {
        kitty := &Cat{"kitty"}
        // kitty.Eat()

        spike := &Dog{"spike"}
        // spike.Eat()

        AnimalEat(kitty)
        AnimalEat(spike)
}
```

Go的多态是通过Interface实现的，我们定义了`Aniaml`的接口，然后`Cat`与`Dog`都实现了它的方法，另外提供了`func AnimalEat(animal Animal)`，在该函数中调用animal的Eat方法，它不需要知道具体的实现类是谁，它只知道animal实例是一个Animal，它肯定拥有Eat方法。

在测试`TestPolymorphism`中，实例化了Cat kitty与Dog spike，将他们传给函数AnimalEat，结果是：

```
Cat kitty is eating.
Dog spike is eating.
```

这便是多态在Go中的体现。

有其他面向对象语言经验的童鞋可能会问，那struct组合struct有多态吗？比如在Java中，子类可以继承自父类，只要不同的子类他们都相同相同的父类，那么便可以将子类参数赋值给父类，在方法中达到多态的效果。

解答这个问题得回顾下struct的特性，struct的组合嵌套，其实并不是「继承」，它们的关系是 **has-a** 而非 **is-a**，并没有父类子类的概念，所以这在Go语言中是行不通的。

