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

### 问题1：为什么不用 binary? 而用 base64?