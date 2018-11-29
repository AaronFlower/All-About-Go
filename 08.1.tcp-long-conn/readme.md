## TCP Long Connection
TCP 长链接。

### Client Test
```
bash -c 'i=0; while [ $i -lt 5 ]; do go run main.go :2048; i=$[$i + 1];done'
server response--->>: Your request has been submited.2018-11-28 14:58:48.131922 +0800 CST m=+3.440911790
server response--->>: Your request has been submited.2018-11-28 14:58:58.65006 +0800 CST m=+13.958737540
server response--->>: Your request has been submited.2018-11-28 14:59:09.114544 +0800 CST m=+24.422910703
server response--->>: Your request has been submited.2018-11-28 14:59:19.544147 +0800 CST m=+34.852203130
server response--->>: Your request has been submited.2018-11-28 14:59:30.123572 +0800 CST m=+45.431314154
```
### Server Output
```
The server is listening: :2048
Client Established, clien say  HEAD / HTTP/1.0

 2018-11-28 14:58:48.13176 +0800 CST m=+3.440749583
err = read tcp 127.0.0.1:2048->127.0.0.1:53300: i/o timeout
Client Established, clien say  HEAD / HTTP/1.0

 2018-11-28 14:58:58.650043 +0800 CST m=+13.958720831
err = read tcp 127.0.0.1:2048->127.0.0.1:53303: i/o timeout
Client Established, clien say  HEAD / HTTP/1.0

 2018-11-28 14:59:09.114528 +0800 CST m=+24.422893888
err = read tcp 127.0.0.1:2048->127.0.0.1:53312: i/o timeout
Client Established, clien say  HEAD / HTTP/1.0

 2018-11-28 14:59:19.544129 +0800 CST m=+34.852185445
err = read tcp 127.0.0.1:2048->127.0.0.1:53315: i/o timeout
Client Established, clien say  HEAD / HTTP/1.0

 2018-11-28 14:59:30.123546 +0800 CST m=+45.431287996
err = read tcp 127.0.0.1:2048->127.0.0.1:53321: i/o timeout
^C
Command terminated

Interrupt: Press ENTER or type command to continue
```
