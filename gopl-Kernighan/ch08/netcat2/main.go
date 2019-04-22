package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

/**
 * 作为客户端程序，一直从 os.Stdin 拷贝数据，然后发送给服务器。而启用的协程是将服务器的响应数据拷贝到
 * os.Stdout 上。
 * 该程序有个问题，即服务器关闭了之后，客户端并不知道去关闭，还可以从 os.Stdin 读取数据。服务器已经关闭
 * 再去发送数据就会导致出错。我们应该当服务器关闭后，我们客户端也自动退出才行。这就需要用到 unbuffered channel
 */
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)

	// io.Copy copies from src to dst until either EOF is reached  on src or an error occurs.
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("The serser has been closed")
		log.Fatal(err)
	}
}
