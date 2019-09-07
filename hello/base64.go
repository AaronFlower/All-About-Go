package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 对 "A" 进行 Base64 编码，其 ascii 编码是 65, 其二进制为 (1000 0011)
	const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	const encodeURL = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	const encodeStr = "zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA0123456789-_"
	msg := "C"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("'C' base64 std encoded with padding: ", encoded)

	encoded = base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(msg))
	fmt.Println("'C' base64 std encoded without padding: ", encoded)

	encoder := base64.NewEncoding(encodeStr)
	encoded = encoder.EncodeToString([]byte(msg))
	fmt.Println("'C' base64 custom encoded with padding: ", encoded)

	encoded = encoder.WithPadding(base64.NoPadding).EncodeToString([]byte(msg))
	fmt.Println("'C' base64 custom encoded without padding: ", encoded)

}
