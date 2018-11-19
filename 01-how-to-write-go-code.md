## Go Tutorial

### 1. 环境配置

```dockerfile
FROM alpine:3.7

RUN echo http://mirrors.ustc.edu.cn/alpine/v3.7/main/ > /etc/apk/repositories
RUN apk update
RUN apk add git zsh vim gcc
RUN sh -c "$(wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O -)" ; exit 0;
# install go
RUN wget https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz -O go.tar.gz
RUN tar -C /usr/local -xzf go.tar.gz
RUN rm -rf go.tar.gz

# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
ENV PATH $PATH:/usr/local/go/bin
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# add go workspace
RUN mkdir $HOME/go

# init vim
# RUN git clone https://github.com/AaronFlower/YAV-dotvim.git ~/.vim
# RUN ln -s ~/.vim/vimrc ~/.vimrc
# RUN vim +PlugInstall +qall
# RUN vim +GoInstallBinaries +qall

CMD ["/bin/zsh"]
```

### 2. 运行

```bash
docker run -it --name=my-alpine-go local/alpine-go
```

### 3. go-tour

```bash
cd ~/go
go get -v golang.org/x/tour
```

### 4. How to write Go Code

如有遗忘，请看[视频](https://www.youtube.com/watch?v=XCsL89YtqCs)。

```bash
# 设置 code 目录
mkdir gocode
export GOPATH=$HOEM/gocode
mkdir -p src/github.com/af
cd src/github.com/af

# 新建项目
mkdir hello
cd hello
vim hello.go

# 构建安装项目
go install # build and install in the bin path
~/gocode/bin/hello

# 可以把 bin 目录设置到系统路径中去。
export PATH=$HOME/gocode/bin:$PATH
hello

# create package, package can be import by other package.
cd af
mkdir string

vim stirng.go
go build

# after go build go install
ls ~/gocode/pkg/darwin_amd64/github.com/nf
string.a # 静态库

# use package
vim hello.go # add string
go install
hello # reversed.

# add test
cd ../string
vim string_test.go
go test #dose handle unicode
# byte-->rune


```

#### string.go

```go
package string

func Reverse(s string) string {
    b := []byte(s)
    for i:=0; i < len(b)/2; ++i {
        j := len(b) - i - 1;
        b[i], b[j] = b[j], b[i]    
    }
    return string(b)
}
```

#### hello.go

```go
package main

import (
    "fmt"
    "github.com/nf/string"
)

func main() {
    fmt.Println(string.Reverse("Hello, new gopher!"))
}
```

#### string_test.go

```go
package string
import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

