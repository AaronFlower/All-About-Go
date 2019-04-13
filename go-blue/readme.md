## Should kown

### Logical Operators: !, &&, ||
### go 标准库: net, http,
### bufio
通过 `bufio.NewScanner()` 创建一个 `bufio.Scanner` 对象，该对象从 `os.Stdin` 中读取内容。

```
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
        // Do your job with s.Text()
		fmt.Println("input:", s.Text())
	}
```

`bufio.Scanner` 其实需要传进来的源是 `io.Reader` 。

### os
- os.Stdin, os.Stdout, os.Stderr

### md5


```
m := md5.New()
io.WriteString(m, strings.ToLower(user.Email()))
```

### bytes
### time

- var buf = bytes.Buffers; buf.String()

### strings
- strings.Split(r.URL.Path, "/")
- strings.Split(r.URL.Path, "/")

### Go-Blue

ch01:
    chat, room, client, tracer
ch02:
    Middleware, wrapper handler
ch03:

ch05:

#### Ch05 Building Distributed System and Working with Flexible Data

##### MongoDB
BSON（/ˈbiːsən/）是一种计算机数据交换格式，主要被用作MongoDB数据库中的数据存储和网络传输格式。它是一种二进制表示形式，能用来表示简单数据结构、关联数组（MongoDB中称为“对象”或“文档”）以及MongoDB中的各种数据类型。BSON之名缘于JSON，含义为Binary JSON（二进制JSON）。

效率：与 JSON 相比，BSON 着眼于提高存储和扫描率。BSON 文档中的大型元素以长度字段为前缀以便于扫描（给出长度当然快了）。在某些情况下，由于长度前缀和显式数组索引的存在，BSON 使用的空间会多于 JSON。

例子：一个内容为 `{"hello": "world"}` 的文档将存储为：

```
Bson:
  \x16\x00\x00\x00               // 总文档大小
  \x02                           // 0x02 = 类型：String（字符串）
  hello\x00                      // 字段名
  \x06\x00\x00\x00world\x00      // 字段值（值大小长度，值，空终止符）
  \x00                           // 0x00 = 类型：EOO（'end of object'，对象结尾）
```

- [BSON, Binary JSON](https://zh.wikipedia.org/wiki/BSON)
- [Mongo Shell Basci Query Documents](https://docs.mongodb.com/manual/tutorial/query-documents/)

### Ch06 Exposing Data and Functionality through a RESTful Data Web Service API

- How to safely share data between HTTP handlers using the `context` package?

#### Sharing data between handlers

有时候我们需要在 Middleware 和 Handlers 之间共享数据。Go 1.7 在标准库引入了 `context` 包，为在请求范围内提供了共享数据的方法。

通过 `request.Context()` 方法可以为每一个 `http.Request` 创建一个 `context.Context` 对象，我们可创建新 context 对象。接下来我们可调用 `request.WithContext()` 方法来获取`http.Request` 的浅拷贝。哈哈，翻译不好...

Every `http.Request` method comes with a `context.Context` object accessible via the `request.Context()` method, from which we can create new context objects. We can then call `request.WithContext()` to get a (cheap) shallow copied `http.Request` method that uses our new Context object.

添加一个值，我们可以创建一个新的 `context` (当然是基于请求的 context 创建）：

```
ctx := context.WithValue(r.Context(), "key", "value")
```

**Note**: 添加的新值只推荐是 Strings, Integers 类型的，不要是指针或其它对象之类的。

在 Middleware 中我可以把创建的新 context 传递给下一个 Handler.

```
next.ServerHTTP(w, r.WithContext(ctx))
```

[context package](https://golang.org/pkg/context/)


#### Context Keys

在 [Context](https://golang.org/pkg/context/) 的源码中我们可看到定义的 WithValue 方法中的 key，value 的类型是 `interface{}` 的，即 key, value 是任意的。

```
type Context interface {
    // Deadline returns the time when work done on behalf of this context
    // should be canceled. Deadline returns ok==false when no deadline is
    // set. Successive calls to Deadline return the same results.
    Deadline() (deadline time.Time, ok bool)

    // a Done channel for cancelation
    Done() <-chan struct{}

    Err() error
}
```
- func WithValue

```
// WithValue returns a copy of parent in which the value associated with key is val.
func WithValue(parent Context, key, val interface{}) Context
```

#### Wrapping handler functions

在第二章中我们用 MustAuth Middleware 的时候写了一个 wrapper 函数，用包裹 handler: wrapping handler. 在这里用同样的策略来包裹 `http.HandlerFunc` (即创建的函数返回 `http.HandleFunc` 即可)

#### API Keys

我们先创建一个 `withAPIKey` 的包裹函数, 为传入的 http.Request 的 context 带上 APIKey，这样就方便多了。

```
func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if !isValidAPIKey(key) {
			responseErr(w, r, http.StatusUnauthorized, "invalid API key")
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyAPIKey, key)
		fn(w, r.WithContext(ctx))
	}
}
```

#### Cross-origin resource sharing, 加入跨域的包裹函数

```
func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}
```

在真正的代码你应该考虑用 [cors package](https://github.com/fasterness/cors)

#### Injecting dependencies, 依赖注入

我们需要考虑的一个问题，我们需要为所有的 handler 都创建一个 DB 链接吗？ DRY (Don't repeat yourself).

实际，我们可以创建一个新类型来我们的 handlers 封装进来必要的依赖对象。现在创建一个新的 Server 类型：

```
// Server is the API server
type Server struct {
	db *mgo.Session
}
```

#### Responding

API 如何影响数据？这里包括状态码，数据，错误信息和一些 headers。借助于`net/http` 包可很方便的完成这些功能。

创建一个 `respond.go` 文件完成我们代码吧。

```
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	// 创建一个 json Decoder 把 r.Body reader 中读取内容解析到 v 中
	return json.NewDecoder(r.Body).Decode(v)
}

func encodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	// 创建一个 json Encoder 把 v 的内容编码到 w writer 中.
	return json.NewEncoder(w).Encode(v)
}

func respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		encodeBody(w, r, data)
	}
}

func respondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func respondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	respondErr(w, r, status, http.StatusText(status))
}
```

#### Understanding the request

`http.Request` 对象为我们提供了 HTTP request 的所有信息，可以去 `net/http` 文件中查看。包含的信息不限如下这些：

- URL, Path, query
- HTTP Method
- Cookies
- Files
- Form Values
- The referrer and uses angent of requester
- Basic authentication details
- The request body
- the header information.

但是 http.Request 不会我们解析 REST path 的路径，如果我们使用 Gorilla 的 mux 会我们提供这些功能，但是我们简单的程序我们可自己实现对 path 的解析，写个解析函数就行了.

#### Serving our API with once function

创建我们的 Server

#### Handling endpoints

定义我们的 `poll` 数据结构 。

#### Complete the handlers

## 7 Random Recommendations Web Service

### The project overview

- 定义一个 `:8082/journeys` API 接口

但是我们不能把数据完全暴露给用户，所以我们用 `Facade` 封装一次。Facade 接口仅提供一个 `Public` 方法返回一个结构体的 view （如果该结构体也实现了 Public 方法的话，否则的则返回该对象）。

```

package meander

// Facade defines a Public method to view a struct view
type Facade interface {
	Public() interface{}
}

// Public returns a struct view
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
```

Normally, single method interface such as our Facade are named after the method they describe, such as Reader and Writer. However, Publicer is just confusing, so I deliderately broke the rule.


The Facade interface and Public method are probably more expressive and give us more control.

### Generating random recommendations

构建查询 Google Places API 的接口。

### Enumerators in GO (enum) 枚举类型

我们使用 `itoa` 来帮我们创建枚举类型，并且把忽略 0 值。

```
package meander

type Cost int8

const (
    _ Cost = iota
    Cost1
    Cost2
    Cost3
    Cost4
    Cost5
)
```
包含一系列的测试。


## ch08 Filesystem Backup

文件系统备份。

创建一个文件备件系统，来备件我们的代码源文件，在我们的项目源代码发生变化的时候进行备份。比如：文件内容更新保存，删除添加文件等。我们可以在任何时间结点上进行恢复。

本章将学到的主题：

- 怎样组织包含 package 和 command-line 工具的项目目录
- 怎样用程序化的方法来持久化简单数据
- os 标准库操作文件系统
- 如何响应 ctrl-c 信号
- `filepat.Walk` 来迭代遍历文件夹及文件
- 如何发现文件是否发生了变化
- 如何用 archive/zip 打包 zip 文件
- 构建工具是怎么处理命令行 flags 的


解决方案设计：

1. 解决方案可以定期根据时间间隔对文件进行备份
2. 可以控制时间间隔
3. 快照文件可以压缩
4. 快速迭代
5. 项目方案易扩展
6. 两个命令行工具：一个是后端守护进程进行完成我们的备份功能，一个是前端交互工具可以让我们向后台守护进程发起查看，添加删除目录的服务请求。


## ch09  Building a Q&A Application for Google App Engine

构建一个类似 StackOverflow, Quora 的问答网站。

本章将学到的主题：

- 利用 Google App Engine SDK for GO 构建测试应用程序并且布署到云端
- 使用 YAML 文件配置应用程序
- 使用 Google App Engine 的 Modules 来组织你的应用
- 使用 Google Cloud Datastore 来持久化数据
- 数据建模
- 使用 Google App Engine 来对用户进行授权验证
- 如何进行反范式数据存储实例(与 RMDBS 的范式相反）
- 怎么复用事务保证数据的完整性
- 怎么提高代码的可维护性
- 如何实现简单的 HTTP 路由

### The Google App Engine SDK for Go

为了运行的布署 Google App Engine 的应用，我们首先需要下载配置其 Go 语言的 SDK 。

### Denormalizing data

Developers with experience of relational databases(RMDBS) will often aim to reduce data redundancy (trying to have each piece of data appear only once in their database) by **normalizing** data, spreading it across many tables, and adding refrences(foreign keys) before joining it back via a query to build a complete picutre.

In schemaless and NoSQL databases, we tend to do the opposite, we **denomralize** data so that each document contains the complete picture it needs, making read times extremely fast since if only to go and get a single thing.

Twitter 是怎么解决 "the Lady Gaga Problem" 问题的那？


### Line of sight in code

The cost of writing a function is relatively low compared to the cost of maintaining if, especially in successful, long-running projects. So it's worthing taking the time to ensure the code is readable by our future selves and others.

- Align the happy path to the left edge so that you can scan down a single column and see the expected flow of execution.
- Don't hide the happy path logic inside a nest of indented braces
- Exit early from your function
- Indent only to handle errors or edge cases
- Extract functions and emthods to keep bodies small and readable

happy path: 如果一切正常的话，函数所执行的代码路径，即正确逻辑。而 非 happy path, 是处理异常，错误发生时才会执行的路径。

### Context in Google App Engine

`Context` is actually an interface that provides **cancelation signals, execute deadlines, and request-scoped data througout a stack of function calls across many components and API boundaries**.
