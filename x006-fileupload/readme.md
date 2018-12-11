## File Upload

The `File` interface doesn't define any methods, but inherits methods from the `Blob` interface:

- `Blob.slice([start[, end[, contentType]]])

### The FileReader API

从 [FileReader API](https://w3c.github.io/FileAPI/#dfn-filereader) 可以出其方法和常量的定义。

```idl
[Constructor, Exposed=(Window,Worker)]
interface FileReader: EventTarget {

  // async read methods
  void readAsArrayBuffer(Blob blob);
  void readAsBinaryString(Blob blob);
  void readAsText(Blob blob, optional DOMString encoding);
  void readAsDataURL(Blob blob);

  void abort();

  // states
  const unsigned short EMPTY = 0;
  const unsigned short LOADING = 1;
  const unsigned short DONE = 2;


  readonly attribute unsigned short readyState;

  // File or Blob data
  readonly attribute (DOMString or ArrayBuffer)? result;

  readonly attribute DOMException? error;

  // event handler content attributes
  attribute EventHandler onloadstart;
  attribute EventHandler onprogress;
  attribute EventHandler onload;
  attribute EventHandler onabort;
  attribute EventHandler onerror;
  attribute EventHandler onloadend;

};
```

- 问题1：为什么不用 binary? 而用 base64?
- 问题2：怎么标识一批请求？怎样避免冲突？

客户端加上时间戳标识, 上传就不会冲突了，就可以惟一标识了。

```
test master ✗ 4h39m △ ➜ ll |awk 'NR > 1 {print $NF}'|xargs -L 1 md5
MD5 (1679d561729eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d561da3eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d565ca8eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d56643feclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d566590eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d566b95eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d566d30eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d56a30beclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d56aac0eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
MD5 (1679d56ad64eclipse-inst-mac64.tar.gz) = dd73186952cbb9960bff92652dfa5f53
```
