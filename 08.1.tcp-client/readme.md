### Client Test
```
bash -c 'i=0; while [ $i -lt 5 ]; do go run main.go :7777; i=$[$i + 1];done'
```
可以看出客户端每次都是通过不同的随机端口来请求服务器。另外从客户端的输出 
`tcp 127.0.0.1:64741->127.0.0.1:7777` 可以看出我们使用的协议是 `tcp`, ip 是本地，
服务器的端口是固定的，而客户端的端口是随机的。
所以网络通信中三元组 `(ip, 协议，端口)`可以标识网络的进程了，网络中利用这个三元组进行交互。

### Client Output
```
err = read tcp 127.0.0.1:64741->127.0.0.1:7777: read: connection reset by peer
exit status 1
err = read tcp 127.0.0.1:64744->127.0.0.1:7777: read: connection reset by peer
exit status 1
err = read tcp 127.0.0.1:64748->127.0.0.1:7777: read: connection reset by peer
exit status 1
err = read tcp 127.0.0.1:64751->127.0.0.1:7777: read: connection reset by peer
exit status 1
err = read tcp 127.0.0.1:64754->127.0.0.1:7777: read: connection reset by peer
exit status 1
```
### Server Output
```
The server is listening: :7777
conn established. 2018-11-28 11:30:39.362185 +0800 CST m=+205.391693079
conn established. 2018-11-28 11:30:39.698409 +0800 CST m=+205.727906360
conn established. 2018-11-28 11:30:40.115138 +0800 CST m=+206.144623480
conn established. 2018-11-28 11:30:40.474373 +0800 CST m=+206.503847535
conn established. 2018-11-28 11:30:40.82431 +0800 CST m=+206.853774291
```
