## Func convertion in Golang

在 Golang 中函数也是一种类型，所以也可以进行类型转换。也可以实现接口定义的方法。
在和接口一起应用时会很强大。

### 例子1 
Go 的 `error` 接口定义如下：

```go
type error interface {
    Error() string
}
```

这就意味着对于任何类型只要实现 `Error () string` 方法都可以被赋值给 `error`.

下面我们定义一个 `binFunc` 类型，并且实现 `error` 的接口。

```go
type binFunc func(int, int) (int, error)

func (f binFunc) Error() string {
    return "binFunc error"
}
```

下面我们定义一个 divide 函数

```go
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Can't be divided by zero")
	}
	return x / y, nil
}
```

就可以把 divide 函数转换成 binFunc 统一处理了。

```go
func main() {
	var bf = binFunc(divide)
	o, err := bf(3, 0)
	if err != nil {
		fmt.Println(bf)
		fmt.Println(err)
	} else {
		fmt.Println("The result is :", o)
	}
}
```


### 例子2

Go `net/http` 中的 Handler 也是一个接口，定义了一个 `ServerHTTP` 方法。

```
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

我们在使用 `http.HandleFunc` 函数对路由进行绑定的时候，也可使用自己的实现的 Handler，如我们可以通过
自己实现的 Hanlder 来统一的处理错误。
如,自己定义一个 `appHandler` 类型。

```go
type appHandler func(http.ResponseWriter, *http.Request) error 

// 实现 ServeHTTP 该函数没有任何返回值。
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        http.Error(w, err.Error(), 500)
    }
}
```

上面我们自定义的路由器，然后我们可以通过如下方式来注册函数：

```go
func init() {
    http.HandleFunc("/view", appHandler(viewRecord))
}

func viewRecord(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return err
	}
	return viewTemplate.Execute(w, record)
}
```

其实我们可以把这个错误信息定义的更加完美：

```go
type appError struct {
    Error error
    Message string
    Code int
}
```
这样我们的自定义路由器可以改成如下方式：

```
type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		c := appengine.NewContext(r)
		c.Errorf("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}
```

这样修改完自定义错误之后，我们的逻辑处理可以改成如下方式：

```go
func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		return &appError{err, "Record not found", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appError{err, "Can't display record", 500}
	}
	return nil
}
```
如上所示，我们访问 view 的时候可以根据不同的情况获取不同的错误码和错误信息，虽然代码量差不多，
但是这个错误更加明显，提示的错误信息更加友好，扩展性也比第一个好。

