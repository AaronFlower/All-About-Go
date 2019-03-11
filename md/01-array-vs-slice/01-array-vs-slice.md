## Array vs Slice

下面的两行代码有什么区别？

```go
a := [3]int{1, 2, 3 }
b := []int{1, 2, 3}
```

区别是 a 是一个 Array, 而 b 是一个 slice.

对于函数：

```go
func walk(values []int) {
	for _, v := range values {
		fmt.Println(v)
	}
}
```

其第一个参数是一个 `slice` 所 `walk(b)` 是正确的调用方法，而 `walk(a)` 则是错误的调用方法。


1. Array 是具有固定长度，赋值操作是拷贝语义的。
2. Slice 不有固定长度，赋值操作是是引用语义的。
3. Slice 其实是下个结构体，有 len, cap 属性。
4. `len`, `cap` 函数对于 Aarry 和 Slice 都是适用的。

创建 Array 的方法：
```go
a := [3]int{1, 2, 3}
b := a
```

创建 Slice 的方法:
```go
a := []int{1, 2, 3}
b := a
c := a[1:3]
d := make([]int, 3)
```


### References

1. [Array vs Slice](https://stackoverflow.com/questions/21719769/passing-an-array-as-an-argument-in-golang/21722697#21722697)
2. [The Minimum You Need To Know About Arrays and Slices in Golang](https://www.openmymind.net/The-Minimum-You-Need-To-Know-About-Arrays-And-Slices-In-Go/)
