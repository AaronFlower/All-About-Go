## Middleware Patterns in Go

A middleware handler is simply an `http.Handler` that wraps another `http.Handler` to do dome pre-and/or post-processing of the request. It's called "middleware" because it sits in the middle between the Go webserver and the actual handler.

### Logging Middleware

下面是我们的服务器程序，未加任何中间件。
```go
  package main

  import (
      "fmt"
      "log"
      "net/http"
      "time"
  )

  func helloHandler(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("Hello world\n"))
  }

  func currentHandler(w http.ResponseWriter, r *http.Request) {
      curTime := time.Now().Format(time.Kitchen)
      w.Write([]byte(fmt.Sprintf("the current time is %s", curTime)))
  }

  func main() {
+     addr := ":8080"
+     fmt.Println("Run....")
      mux := http.NewServeMux()
      mux.HandleFunc("/v1/hello", helloHandler)
      mux.HandleFunc("/v1/time", currentHandler)

~     err := http.ListenAndServe(addr, mux)
      if err != nil {
          log.Fatal(err)
      }
~     fmt.Printf("Server is listening at %s", addr)
+     log.Printf("Server is listening at %s", addr)
  }
```

执行结果 
```bash
x004-middleware master ✗ 2m △ ◒ ⍉ ➜ http :8080/v1/time
HTTP/1.1 200 OK
Content-Length: 26
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Dec 2018 13:15:11 GMT

the current time is 9:15PM

x004-middleware master ✗ 2m △ ◒ ➜ http :8080/v1/hello
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Dec 2018 13:15:17 GMT

Hello world

```

下面我们相为所有的请求添加一个 log 记录。我们可以为所有的 handle 都写上同样的记录代码，但是最好的方法是在一个地方统一处理就好。

通过定义一个新的结构体实现  `http.Handler` 的 `ServeHTTP()`方法即可。这个结构体中有一个 `handler` 成员变量用来存储请求时真正的 `http.Handler`。

所以该结构体的定义如下：
```go
+ // Logger is a middleware handler that does request logging
+ type Logger struct {
+     handler http.Handler
+ }
+
+ // ServeHTTP handles the request by passing it to the real
+ // handler and logging the request details.
+ func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
+     start := time.Now()
+     l.handler.ServeHTTP(w, r)
+     log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
+ }
+
+ // NewLogger constructs a new Logger middleware handler
+ func NewLogger(handlerToWrap http.Handler) *Logger {
+     return &Logger{handlerToWrap}
+ }
+
```

`NewLogger()` 构造函数接收一个 `http.Handler` 参数返回一个 Logger 的实例。因为 `http.ServeMux`实现了 `http.Handler`接口， 所以可以把这个 `mux`传递给 logger 中间件。另外，因为我们 Logger 也实现了 `ServeHTTP()`方法，所以实现了 `http.Handler`接口，所以我们的 Logger 实例也可以传递给 `http.ListenAndServe()`函数。这们我们的 Logger 中间件就应用上了。

通过请求可看到服务器的 Log 信息:
```bash
Interrupt: Press ENTER or type command to continue
Run....
2018/12/05 21:41:30 GET /v1/hello 12.982µs
2018/12/05 21:41:39 GET /v1/hello 8.136µs
2018/12/05 21:41:43 GET /v1/time 22.956µs
2018/12/05 21:41:50 GET /v1/time 14.023µs
2018/12/05 21:41:51 GET /v1/time 11.245µs
2018/12/05 21:41:52 GET /v1/time 10.273µs
```

### Chaning Middleware

因为中间件的构造函数接收的参数和返回都是 `http.Handler`，所以你可以对多个中间件进行链式调用。

下面的例子，我们想为所有的 handler 添加 header 返回信息。我们可以创建下面的中间件。

```go
+ // ResponseHeader is a middleware handler that adds a header to the response.
+ type ResponseHeader struct {
+     handler     http.Handler
+     headerName  string
+     headerValue string
+ }
+
+ // ServeHTTP handles the request by adding the response header
+ func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
+     // add the header
+     w.Header().Add(rh.headerName, rh.headerValue)
+
+     // call the wrapped handler
+     rh.handler.ServeHTTP(w, r)
+ }
+
+ // NewResponseHeader constructs a new ResponseHeader middleware handler
+ func NewResponseHeader(handler http.Handler, headerName string, headerValue string) *ResponseHeader {
+     return &ResponseHeader{
+         handler,
+         headerName,
+         headerValue,
+     }
+ }
+
```
修改我们 `mux`：
```go
    wrappedMux := NewLogger(NewResponseHeader(mux, "Midde-Ware-Header", "Foo Value"))
```
请求结果可以看到中间件输出的 heade 信息。

```bash
x004-middleware master ✗ 1m △ ◒ ➜ http :8080/v1/time
HTTP/1.1 200 OK
Content-Length: 26
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Dec 2018 13:57:22 GMT
Midde-Ware-Header: Foo Value

the current time is 9:57PM

x004-middleware master ✗ 14m △ ◒ ➜ http :8080/v1/hello
HTTP/1.1 200 OK
Content-Length: 12
Content-Type: text/plain; charset=utf-8
Date: Wed, 05 Dec 2018 13:57:31 GMT
Midde-Ware-Header: Foo Value

Hello world

```

你可以根据需要添加更多的中间件。当中间件很少的时候，这样写没有问题，但是如果有更多的中间件，你可能需要 [Adapter Pattern](https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81）, 这样可以更加优雅把你的中间件组合在一起。

### Middleware and Request-Scoped Values

