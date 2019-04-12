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
    Middleware
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
