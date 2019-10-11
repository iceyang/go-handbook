# 数组和切片

`数组 array`是许多语言中都会有的基本数据结构，它是我们开辟的存储相同数据类型的一连串地址空间（长度不可变），通过下标寻址，随机访问性能为O(1)。

`切片 slice`可以当做是`长度可变`的数组，在使用上与数组相同（初始化的方式有区别）。

## 数组(array)

我们可以通过下面的方式初始化数组，下面的例子都是初始化长度为5的数组：

```Go
// 声明式
var arr1 [5]int

// 直接赋值
arr2 := [5]int{}

// 赋值并初始化值
arr3 := [5]int{1,2,3,4,5}

// 赋值，部分初始化
arr4 := [5]int{1,2,3}

// 赋值，自动计算数组长度
arr5 := [...]int{1,2,3,4,5}

// 赋值，通过下标初始化
arr6 := [5]string{0: "first", 2: "third", 1: "second"}
```

> 注意，数组在声明或赋值初始化后，对于没有主动初始化的空间，它会使用被声明的类型的初始值进行填充。
> 比如上面的arr4，它的具体内容将是：[1 2 3 0 0]

在数组的使用上，我们通过下标进行访问，下标从`0`开始：

```Go
var arr7 [5]int 
// arr7: [0 0 0 0 0]
arr7[0] = 1
// arr7: [1 0 0 0 0]
arr7[4] = 1
// arr7: [1 0 0 0 1]
```

## 切片(slice)

切片在初始化上与数组不太相同，我们不需要指定长度，另外可以使用内建函数`make`来开辟空间：

`下面会提到 *容量* 这个上面没有提到的词，可先暂且知道它的存在，更后面会详细说到`

```Go
// 声明式，这是长度为0的切片
var slice1 []int

// 赋值初始化，切片长度为5，容量为5
slice2 := []int{1, 2, 3, 4, 5}

// 通过make创建长度为5，容量为5的切片
slice3 := make([]int, 5)

// 通过make创建长度为5，容量为10的切片
slice4 := make([]int, 5, 10)

// 通过make创建长度为0，容量为10的切片
slice5 := make([]int, 0, 10)

// 赋值，通过下标初始化
slice6 := []string{0: "first", 2: "third", 1: "second"}
```

> 跟数组相似，当我们为切片开辟了空间之后，没有主动初始化的空间，将会使用被声明的类型的初始值填充。

同样的，切片通过下标来访问切片元素，

```Go
slice7 := make([]int, 5)
// slice7: [0 0 0 0 0]
slice7[1] = 2
// slice7: [0 2 0 0 0]
slice7[3] = 3
// slice7: [0 2 0 3 0]
```

### 切片的扩容

内置函数`append`可以对切片进行元素的添加。这也是切片跟数组最大的区别之处。

```Go
slice8 := []int{1, 2, 3, 4, 5}
// 追加单个元素
slice8 = append(slice8, 6)
// 追加多个元素
slice8 = append(slice8, []int{7, 8}...)
slice8 = append(slice8, 9, 10)
```

## 截取片段

数组和切片支持获取片段的功能：

```Go
arr8 := [5]int{1, 2, 3, 4, 5}
slice9 := []int{1, 2, 3, 4, 5}

// 从第二个元素开始截取
arr8[1:]
// 开始截取到第3个元素
arr8[:3]
// 截取下标[1,3)的元素
slice9[1:3]
```

> 需要注意的是，截取是左闭右开区间。

## 长度与容量

数组和切片，它们都有`长度`和`容量`的概念，只不过对于数组来说，容量便等于长度。

对于切片来说，`容量`是针对其底层数组而言的，当我们使用`append`对切片进行增加的时候，假如增加后的切片长度超过了容量，那么切片需要对底层数组进行`扩容`，这个扩容，实际上是重新开辟一片数组空间，同时给到`新`的切片。如果增加后的长度没有超过容量，那么返回的切片还是原来的底层数组，不会有所改变，只是看到的内容变多了。

> 所以在使用`append`时，其实是隐藏了一些`陷阱`在里面，文章的最后面会对这部分继续描述。

通过内置函数`len`，我们可以得到数组和切片的实际长度，而通过内置函数`cap`，我们可以得到数组和切片的容量：

```Go
slice10 := make([]int, 5, 10)
length := len(slice10)
capacity := cap(slice10)
```

## 遍历

通过关键字`range`，我们可以很方便地对数组或切片进行遍历：

```Go
arr9 := [...]int{1, 2, 3, 4, 5}
for index, value := range arr9 {
	fmt.Printf("index: %d, value: %d\n", index, value)
}
```

## 越界错误

数组和切片的访问，下标是从0开始，假设访问的下标>=长度时，就会发生越界错误。

当对数组的访问越界，而且编译器能直接推断出越界时，编译会不通过，而避免运行时才报错。

```Go
arr10 := [5]int{1, 2, 3, 4, 5}
// _ = arr10[6] // 编译不通过

index := 6
_ = arr10[index] 
```

## 复杂类型

数组和切片的类型，可以是复杂的类型，比如数组、结构体、切片。

当类型是数组时，此时就变成了一个二维数组，或者是更高维的数组：

```Go
// 初始化了一个二维数组
arr11 := [5][5]int{}
// 将第三行第三列赋值为1
arr11[2][2] = 1
// [ [0 0 0 0 0]
//   [0 0 0 0 0]
//   [0 0 1 0 0]
//   [0 0 0 0 0]
//   [0 0 0 0 0] ]
```

类型为结构体时：

```Go
type Student struct {
	Name string
	No   int
}

slice11 := []*Student{
	&Student{"Justin", 1},
}
```

## 切片的底层数组

我们知道，在使用内建函数`append`对切片进行增加时，如果容量不够，则会开辟一片新的内部数组空间，交给新的切片，这个时候我们只需要接收到新的切片进行后续操作即可。而当容量足够时，`append`会直接使用原有空间进行操作。

所以，我们了解到这个细节后，需要留意出现的状况：

* 容量足够时，切片会使用原有内置数组

```Go
slice12 := make([]int, 5, 10)
slice13 := append(slice12, 6)
slice13[0] = 1
// 猜猜这里会输出什么
fmt.Println("slice12: ", slice12)

slice14 := make([]int, 5, 10)
slice15 := append(slice14, 6, 7, 8, 9, 10, 11)
slice15[0] = 1
// 猜猜这里会输出什么
fmt.Println("slice14: ", slice14)
```

我建议你手动运行下结果，然后根据容量与长度的问题稍微思考下，可能会更深刻一些。

* 根据容量，主动为切片设置长度

因为我们知道了切片的底层是个数组后，我们可以手动对切片的长度进行操作，比如：

```Go
slice16 := make([]int, 5, 10)
fmt.Println("slice16: ", slice16)

slice17 := slice16[:10]
fmt.Println("slice17: ", slice17)

// slice18 := slice16[:11] // 这里会产生越界错误
```

## 总结

文章从初始化、使用方式、细节等方面介绍了数组和切片，如有错漏或补充，十分欢迎各位进行补充改正。

本文的具体代码可以在 [此处](../code/array_slice) 获得。