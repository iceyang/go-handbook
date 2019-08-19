# 使用gin搭建HTTP服务

[gin](https://github.com/gin-gonic/gin) 是用`Go`语言编写的一个高性能HTTP Web框架，通过它，我们可以快速搭建一个HTTP服务。

本文并不是一篇gin的入门教程，而是介绍如何基于gin搭建一个完善的web服务，包括代码结构的分层，接口输出的标准化，如何进行错误处理，以及使用时的注意事项。

## 先把服务跑起来

首先，我们先让服务可以运行起来。创建一个项目，编写`main.go`，内容如下：

```Golang
func main() {
        engine := gin.Default()
        engine.GET("/ping", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "message": "pong",
                })
        })
        engine.Run()
}
```

紧接着，运行起来，访问`http://localhost:8080/ping`，可以看到接口返回了`{"message": "pong"}`。

## 定义项目结构

> 关于Go项目的标准化，可以参考[Standard Go Project Layout](https://github.com/golang-standards/project-layout)

我们对项目结构做分层如下：

```
project
├── cmd
├── configs
├── doc
├── internal
├── pkg
│   ├── ctrl
│   ├── model
│   ├── db
│   ├── router
│   └── svc
└── vendor
```

假如我们不希望代码被其他项目使用，可以将代码从`pkg`移到`internal`中。

* cmd：存放的是程序的入口，项目有时不止一个入口，那就可以在这里增加。
* configs：存放配置文件，比如数据库配置。
* doc：程序相关文档
* internal：项目代码(不对外公开)
* pkg：项目代码
    * ctrl：存放控制器
    * model：存放对象模型
    * db：存放数据库连接
    * router：存放路由器
    * svc：service，存放业务逻辑
* vendor: 第三方库

### 更改程序入口

对应回我们前面的程序，我们进行改进。
首先，我们将`目录`建立起来，然后在`cmd`目录下建立`server`目录，将开始时的`main.go`改名`server.go`放入其中：

```
cmd
└── server
    └── server.go
```
以后的程序入口都由cmd包负责。

### 修改路由

在`pkg/router`下面，我们创建文件`engine.go`，内容如下：

```Golang
func Default() *gin.Engine {
        engine := gin.Default()
        engine.GET("/ping", func(c *gin.Context) {
                c.JSON(200, gin.H{
                "message": "pong",
                })
        })
        return engine
}
```

然后修改入口文件`cmd/server/server.go`的内容：

```Golang
func main() {
        engine := router.Default()
        engine.Run()
}
```

以后的接口定义，都在`router`层下编写。

### 新增控制器

接下来，我们需要将控制器提取出来。
在`pkg/ctrl`目录下，我们新增文件`example.go`：

```Golang
type ExampleController struct{}

var Example = &ExampleController{}

func (ec *ExampleController) Ping(c *gin.Context) {
        c.JSON(200, gin.H{
                "message": "pong",
        })
}
```

编辑`pkg/router/engine.go`，修改控制器指向：

```Golang
func Default() *gin.Engine {
        engine := gin.Default()
        engine.GET("/ping", ctrl.Example.Ping)

        return engine
}
```

### 项目结构完成

按照上面的操作，我们已经定义好了项目的结构，业务代码只要按照上面的结构进行存放即可，如有需要，也可以灵活变动。
目前改造完的目录代码结构如下：

```
project
├── cmd
│   └── server
│       └── server.go
├── configs
├── doc
├── internal
├── pkg
│   ├── ctrl
│   │   └── example.go
│   ├── model
│   ├── db
│   ├── router
│   │   └── engine.go
│   └── svc
└── vendor
```

## 定制中间件
### 约定输出格式

后端给到前端的接口内容，我们希望格式是比较标准化的。
比如当接口返回200状态码却不带body时，我们默认输出`{"message": "ok"}`，
当接口返回404却不带body时，我们默认输出`{"message": "找不到资源"}`

按照这样的思维，我们新建文件`pkg/router/middleware.go`：

```Golang
func responseHandler(c *gin.Context) {
        c.Next()
        if c.Writer.Status() == http.StatusNotFound && c.Writer.Size() <= 0 {
                c.JSON(http.StatusNotFound, gin.H{
                        "message": "找不到资源",
                })
                return
        }
        if c.Writer.Status() == http.StatusOK && c.Writer.Size() <= 0 {
                c.JSON(http.StatusOK, gin.H{
                        "message": "ok",
                })
                return
        }
}
```

我们定义了一个`responseHandler`的`gin`中间件，然后修改`pkg/router/engine.go`，配置上中间件：

```Golang
func Default() *gin.Engine {
        engine := gin.Default()
        engine.Use(responseHandler)

        engine.GET("/ping", ctrl.Example.Ping)
        return engine
}
```

为了测试效果，我们在`pkg/ctrl/example.go`中加上控制器对应的内容：

```Golang
func (ec *ExampleController) NotFound(c *gin.Context) {
        c.Status(http.StatusNotFound)
}

func (ec *ExampleController) OK(c *gin.Context) {
        c.Status(http.StatusOK)
}
```

并配置到路由中：

```Golang
engine.GET("/ping", ctrl.Example.Ping)
engine.GET("/404", ctrl.Example.NotFound)
engine.GET("/ok", ctrl.Example.OK)
```

程序执行起来访问对应的接口，你应该能够看到效果了。

### 标准化错误处理

当我们通过`gin.Default()`得到一个引擎实例时，它默认带了错误处理功能(recovery)。
假设我们希望有自己的错误处理流程，比如程序出错时我们需要将错误栈输出到`ELK`之类地方，方便我们进行错误的定位，那我们可以自定义一个错误处理中间件。

下面是一个简单的示例，在实际使用中，如果需要复杂的功能，可以在进行对应的改造。

新建文件`pkg/error.go`，增加错误，并提供相应的`panic`方法：
```Golang
type Error struct {
        Err  error
        Msg  string
        Code int
}

// 将错误直接抛出
func PanicIfErr(err error) {
        if err != nil {
                PanicError(http.StatusInternalServerError, err)
        }
}

// 抛出错误时，携带状态码
func PanicError(code int, err error) {
        panic(&Error{
                Err:  err,
                Msg:  "请求出错，请稍后尝试",
                Code: code,
        })
}

// 自定义状态码和错误信息的错误
func PanicErrorWithMsg(code int, msg string) {
        panic(&Error{
                Err:  errors.New(msg),
                Msg:  msg,
                Code: code,
        })
}
```

在`pkg/router/middleware.go`中，我们加上错误处理的中间件：

```Golang
func recovery(c *gin.Context) {
        defer func() {
                if err := recover(); err != nil {
                        buf := make([]byte, 65536)
                        buf = buf[:runtime.Stack(buf, false)]
                        e, ok := err.(*pkg.Error)
                        if ok {
                                if e.Code >= 500 {
                                        // 错误发送到kibana
                                        log.Printf("%s\n%s", err, buf)
                                }
                                c.AbortWithStatusJSON(e.Code, gin.H{
                                        "message": e.Msg,
                                })
                                return
                        }
                        // 错误发送到kibana
                        log.Printf("%s\n%s", err, buf)
                        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                                "message": "服务出错，请稍后尝试",
                        })
                }
        }()
        c.Next()
}
```

最后，我们不使用`gin`默认的中间件，将我们自己的中间件添加到路由中：

```Golang
func Default() *gin.Engine {
        engine := gin.New()
        engine.Use(recovery)
        engine.Use(responseHandler)

        engine.GET("/ping", ctrl.Example.Ping)
        engine.GET("/404", ctrl.Example.NotFound)
        engine.GET("/ok", ctrl.Example.OK)
        return engine
}
```

*经过这一番改造后，我们对错误处理的规则如下*：

* 当程序发生意料之外的错误时，比如数据库访问出错，我们直接使用`pkg.PanicErr(err)`，将原始error抛出，中间件会处理成500错误，并提示“服务出错，请稍后尝试”，同时将信息记录到日志中；
* 当程序发生在意料之中时：
  * 通过`pkg.PanicErrorWithMsg(code, msg)`，自定义error的msg，这类适用于4xx的客户端错误，比如参数出错，账号不正确，没有权限等等；
  * 通过`pkg.PanicError(code, err)`，将原始错误抛出，同时自定义我们想要的HTTP code。

> 至此，我们应该完成了一个比较简单而统一的web项目结构，这是我对于Go项目的一些应用思考，可能并不适用于所有人，希望能帮到你。

## 其他注意事项
### 路由冲突

现在的`API`设计中，`RESTful`架构大家怕是耳闻能熟了，应用HTTP方法和地址的巧妙设计，使用者能直接猜出接口的作用（语义化）。

当我们在使用`gin`路由的时候，可能会遇到路由设计上的冲突问题（更深层的原因希望以后抽出时间来专门写一篇介绍），这里说明什么情况下会产生路由冲突。

当我们使用`wildcard`参数时，就可能产生冲突，比如：

```
GET /api/users/:id
GET /api/users/mobile/:mobile
```

这两条配置实际上是冲突的，因为后者已经被前者所包括。这里的冲突有两个条件：

1. 一样的HTTP方法
2. 某路由通配符覆盖了其他路由

> 这种冲突在其他语言和框架可能并不会出现，这个时候，我们只能对API进行调整。
