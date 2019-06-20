## Handlers all the way down

在 web 程序中通要写一个中间件来处理特定的逻辑，如验证授权等。在 Golang 中开发中间件 Middleware 的一般模式，先创建创建一个对象，该对象实现 `http.Hanlder` 的 `ServeHTTP` 方法，在 `ServeHTTP` 方法中根据判断决定是否向下继续调用下一层的 handler。我们这个新的对象具有一个 next 字段，用于存储下一个 handler。

另外，我们还需要一个包裹方法，初始化返回我们的中间件 Handler.


例子：比如我们先创建一个 `authHandler`, 然后创建一个 `MustAuth` 包裹方法。

```
package main

import "net/http"

type authHandler struct {
    next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // check code
    h.next.ServeHTTP(w, r)
}

// MustAuth wraps a auth check handler for the next handler.
func MustAuth(next http.Handler) http.Handler {
    return &authHandler{
        next: next,
    }
}

```

Usage: 使用方法

```
http.Handle("/chat", MustAuth(&templHandler{filename: "chat.html"}))
```
