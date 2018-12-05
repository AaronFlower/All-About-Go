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

你可以根据需要添加更多的中间件。当中间件很少的时候，这样写没有问题，但是如果有更多的中间件，
你可能需要 [Adapter Pattern](https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81), 这样可以更加优雅把你的中间件组合在一起。

### Middleware and Request-Scoped Values
现在让我们考虑一个复杂的场景。所设我有几个 handler 都需要验证用户的权限，并且我们有一个函数可以从 `http.Requst`中验证是否是有
权限的权限。如:
```go
func GetAuthenticatedUser(r *http.Request) (*User, error) {
    // Validate the session token in the request,
    // fetch the session state from the session store,
    // and return the authenticated user
    // or an error if the user is not authenticated
}

func UsersMeHandler(w http.ResponseWriter, r *http.Request) {
    user, err := GetAuthenticatedUser(r)
    if err != nil {
        http.Error(w, "please sign-in", http.StatusUnauthorized)
        return
    }

    //GET = respond with current user's profile
    //PATCH = update current user's profile
}

```
我们的 `UsersMeHandler`需要验证用户，所以它会调用 `GetAuthenticatedUser()`来验证是否成功。这们做很正确。
但是如我们其它的 Handlers 也都需要验证用户，如果把相同的代码拷都拷贝到各个 handler 上的话，那就不是很好的实现方法了。实际上用户验证我们也可能通过一个中间件来进行验证。

我们可以像上面定义的中间件一样来定义下面的中间件:
```go
type EnsureAuth struct {
    handler http.Handler
}

func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Response) {
    user, err := GetAuthenticatedUser(r)
    if err != nil {
        http.Error(w, "please sign-in", http.StatusUnauthorized)
        return
    }
    // TODO: call the real handler, but how do we share the user?
    ea.handler.ServeHTTP(w, r)
}

func NewEnsureAuth(handler http.Handler) *EnsureAuth {
    return &EnsureAuth{ handler }
}
```

这里有一个问题？我们获取到了 `user` 信息，那么我们怎么共享 `user`的信息那？这里说的共享不是所有请求的 handler 之间的共享，而是一个请求链下的怎么共享。

这里就要引出 Request Context 了。这是 Go 1.7 引入的。Request Context 可以存储键值对。因为第一个请求的开始的
时候都会实例一个对象，我们可以把我们要共享的信息存入到这个请求对象中。

我们首先定义一个已验证用户的类型和值:

```go
type contextKey int
const authenticateUserKey contextKey = 0
```

然后就可以在中间健的 `ServeHTTP()` 方法中将验证用户信息添加到 Request Context 中了。

```go
func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Response) {
    user, err := GetAuthenticatedUser(r)
    if err != nil {
        http.Error(w, "please sign-in", http.StatusUnauthorized)
        return
    }
    // create a new request context containing the authenticated user
    ctxWithUser := context.WithValue(r.Context(), authenticatedUserKey, user)
    // create a new request using that new context
    rWithUser := r.WithContext(ctxWithUser)
    // call the real handler, passing the new requst
    ea.handler.ServeHTTP(w, rWithUser)
}
```

这样我就把验证的用户信息传递给下一个 handler 了。而且保证之前的中间件不会看到该共享信息。下一个 Handler 
可以通过下面的方式来访问我们的值:
```go
func UsersMeHandler(w http.ResponseWriter, r *http.Request) {
    // get the authenticated user from the request context
    user := r.Context().Value(authenticatedUserKey).(*User)
    // do stuff with that user.
}
```
上面的代码中 `.(*User)`是 type assertion 的使用。

### An Alternative
一些程序员感觉上面的写法还是很不明显，容易使人费解。推荐使用改变函数的签名来实现 request-scoped 变量的传递。
如果 handler 真的需要这些变量，就应该明确声明。

如我们上面的`UsersMeHandler()`会被改写成:

```go
func UsersMeHandler(w http.ReponseWriter, r *http.Request, user *User) {
    // do stuff with the `user`
}
```

改了 Handler 的函数签名就与 HTTP handler 的签名一致了., 所以所需要改一下中间件来适配 http.ServeMux. 修改如下：
```
// AuthenticateHandler is a handler function that also requires a user
type AuthenticateHandler func(http.ResponseWriter, *http.Request, *User)

type EnsureAuth struct {
    handler AuthenticatedHandler
}

func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    user, err := GetAuthenticatedUser(r)
    if err != nil {
        http.Error(w, "please sign-in", http.StatusUnauthorized)
        return
    }

    ea.handler(w, r, user)
}

func NewEnsureAuth(handlerToWrap AuthenticatedHandler) *EnsureAuth {
    return &EnsureAuth{handlerToWrap}
}
```

上面我们构建函数传入的参数是 `AuthenticatedHandler` 类型了。在使用的时候方法如下：
```go
mux.Handle("/v1/users/", NewEnsureAuth(UsersHandler))
mux.Handle("/v1/users/me", NewEnsureAuth(UsersMeHandler))
```
缺点是不在像之前只需要 wrap 下  mux 就行了，现在要 wrap 所有 hanlder 了。解决这个缺点你可以创建一个 
`AuthenticatedServerMux` 类型，接收 `AuthenticatedHandler`方法而不是原本的 `http.Handler`。然后把这个新 mux 回到 main mux 上。

