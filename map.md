# 字典 map
`字典 map`存储的是`key-value`的集合，是非常高效的数据结构。

Go的map使用很简单，声明的时候需要指定`key`的类型、`value`的类型。

定义方式为：map[key_type]value_type

```Go
// 声明式，定义了一个键为string，值为int
var map1 map[string]int

// 赋值声明
map2 := map[string]string{}
```

## 设值

对map的赋值，可以直接用`=`：

```Go
map2["name"] = "Bob"
```

## 取值

从map中取值，也是使用`=`，但是要注意的是，键不存在的情况，需要进行判断处理。

```Go
val, ok := map2["name2"]
```

在上面的例子中，因为name2这个key在map2中不存在，所以我们取到的`val`是空的，在某些场景下，如果没有对是否存在进行校验，那么可能存在预期外的异常或逻辑错误。

`ok`返回的是布尔值，表示是否真正取得了数据。

## 使用range遍历

使用关键字`range`，我们可以对map进行遍历：

```Go
map4 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
}

for k, v := range map4 {
        fmt.Println(k, v)
}
```

如果你运行了上面这段代码，你会发现输出的顺序不是a,b,c。其实是因为，在map的实现里，key是hash的结构存储，而不是顺序的，这也是它存储高效的原因。
