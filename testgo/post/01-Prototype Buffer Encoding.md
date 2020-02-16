## Encoding

Prototype Buffer 在数据序列化时，其表示会远比 JSON 或数据本身要小的多。这是因为 PB 使用一种巧妙的数据协议。但是如果 Message 定义不当的话，空间非旦降低还会增加很多。

### TL;DR

1. 如果 value 有可能为负数，那么就使用 `sint32, sint64`。
2. 如果数字很大时起始 (2^28 = 268435456) 时。
3. 定向 Message 中，常设置的 key-pair 尽量定义在前 31 个 key-pair 中。
4. 如果序列化的数据文件超过 1 MB 时，最好对序列化的数据结构进行下拆分。

### 1. A Simple Message 

假设下面的 message 定义，

```proto
syntax = "proto3";

message MsgInt32 {
    int32 value = 1;
}
```

在自己的应用程序中，

```
	mInt32 := encoding.MsgInt32{
		Value: 150,
	}
```

如果对上面的数据进行序列化，中会占用 3 个字节：

```
08 96 01
```

一个占四个字节的数字被序列化成了 3 个字节。 这是怎么做到的那？

### 1.1 Base 128 Varints

Pb 利用的是 128-varints, 基于 128 的可变 int 表示法。其特点是小的数字占用更少的字节位。
在 varint 中, 每一个字节的最高有效位 (Most Significant Bit, MSB) 来于标识是否还有更多的更多的字节。
所以每个字节只使用 7 位来存储数据信息，从左到右以小字节开始存储。存储的是原数据的补码。

所以，对于数字 1 的表示只需要用一个字节：

```
0000 0001
```

对于 300 则需要两个字节，其表示为：

```
1010 1100 0000 0010
```

将上面的数字转换成 300 的过程如下，1）首先，去掉 MSB 标识位；2）调整顺序。

```
000 0010  010 1100
→  000 0010 ++ 010 1100  // 去掉 MSB 标识位；
→  000 0010 010 1100     // 调整顺序
→  256 + 32 + 8 + 4 = 300
```

总结：
1. 小数占用字节数更小；
2. 无论是 32 位，还是 64 位的负数都占用 10 字节。 (负数被当成的很大的无符号数)

下面是对 `Value` 取不同值时的表示方法。其中 `08` 是 Value Field 的 payload (后面介绍).

```
              -2: 08 fe ff ff ff ff ff ff ff ff 01 
              -1: 08 ff ff ff ff ff ff ff ff ff 01 
               0: 
               1: 08 01 
               2: 08 02 
               3: 08 03 
               4: 08 04 
             150: 08 96 01 
       268435455: 08 ff ff ff 7f 
       268435456: 08 80 80 80 80 01 
      2147483647: 08 ff ff ff ff 07 
     -2147483648: 08 80 80 80 80 f8 ff ff ff ff 01 
```

### 1.2 Message Structure

Pb 定义的 Message 结构其实就是一组 key-value 组合。第一组都有一个序号，这个序号很重要，因为在进行数据序列时，会将该序号作为 key 的一部分存储下来。在编码时，key 与 value 是串在一起的。当解码时如果遇到了不认识的序号，就会跳过。

Key 上承载了两个信息：1）key 的序号; 2) key 的类型（即后面的值是什么类型）。许多语言在实现时，会将这样的 key 称为 tag.

目前 PB 各数据类型的 wire type 定义如下：

| Type | Meaning          | Used For                                                 |
| :--- | :--------------- | :------------------------------------------------------- |
| 0    | Varint           | int32, int64, uint32, uint64, sint32, sint64, bool, enum |
| 1    | 64-bit           | fixed64, sfixed64, double                                |
| 2    | Length-delimited | string, bytes, embedded messages, packed repeated fields |
| 3    | Start group      | groups (deprecated)                                      |
| 4    | End group        | groups (deprecated)                                      |
| 5    | 32-bit           | fixed32, sfixed32, float                                 |

每个 Key 在序列时按 `(field_num << 3) | wire_type` 来进行编码。即类型的 wire_type 占用 3 位（最后三位）。

所以上面的 `08` 对应的二进程 (0000 1000), 即表示序号为 1，数据类型为 varint 的 key。

另外，由于 key 的最后 3 位留给了 wire type, 所以 field number 仅剩下了 5 位。2^5 - 1= 31, 所以对于频次出现高的 key-pair 组尽量放到前 31 个序号中。

### 1.2 ZigZag encoding, Signed Integers

对于无符号数，可以使用 `int32, int64`，但是如果我们的数据中出现出负数，那么就可以考虑有 `sint32, sint64` 数据类型了。这类数据类型使用 ZigZag 来表示数值，也能保证较小的正负数更用更少的字节。

| Signed Original | Encoded As |
| :-------------- | :--------- |
| 0               | 0          |
| -1              | 1          |
| 1               | 2          |
| -2              | 3          |
| 2147483647      | 4294967294 |
| -2147483648     | 4294967295 |

下面是对 `sint32` `Value` 取不同值时的表示方法。

```
               0: 
              -1: 08 01 
               1: 08 02 
              -2: 08 03 
               2: 08 04 
               3: 08 06 
               4: 08 08 
               5: 08 0a 
             150: 08 ac 02 
             300: 08 d8 04 
     -2147483648: 08 ff ff ff ff 0f 
```

即，对于 32-bit 数按： `(n << 1) ^ (n >> 31)` 编码；而对于 64-bit 数按: `(n << 1) ^ (n >> 63)` 来编码。

### 1.3 Non-varint Numbers

对于 `double, sfixed64, fixed64` 其 wire type 为 1，`float, fixed32, sfixed32` 其 wire type 为 5.

### 1.4 Strings

Wire type 为 2 的是 String 类型，在序列化时会存储字符串的长度。下面是不同长度的字符串的二进制表示。

```
               h:0a 01 68 
              he:0a 02 68 65 
             hel:0a 03 68 65 6c 
            hell:0a 04 68 65 6c 6c 
           hello:0a 05 68 65 6c 6c 6f 
     hello world:0a 0b 68 65 6c 6c 6f 20 77 6f 72 6c 64 
```

### 1.5 Embedded Messages

对于嵌入的 Message 类型其 wire type 为 2， 其它表示和正常类型一样。