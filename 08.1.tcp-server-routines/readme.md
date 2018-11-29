### Go Routines
通过使用 goroutine 可以实现服务器同时支持多个并发请求。

### Client Test
```
bash -c 'i=0; while [ $i -lt 20 ]; do go run main.go :1200; i=$[$i + 1];done'
```

### Server Output
可以看到 Server 在秒中多次的响应。只是怎样让服务器发的更频繁那？
```
Client established: 2018-11-28 11:54:34.088372 +0800 CST m=+3.920773497
Client established: 2018-11-28 11:54:34.542266 +0800 CST m=+4.374654396
Client established: 2018-11-28 11:54:34.913353 +0800 CST m=+4.745729965
Client established: 2018-11-28 11:54:35.32417 +0800 CST m=+5.156534144
Client established: 2018-11-28 11:54:35.779241 +0800 CST m=+5.611591503
Client established: 2018-11-28 11:54:36.214168 +0800 CST m=+6.046505753
Client established: 2018-11-28 11:54:36.582933 +0800 CST m=+6.415259712
Client established: 2018-11-28 11:54:36.982568 +0800 CST m=+6.814882750
Client established: 2018-11-28 11:54:37.459202 +0800 CST m=+7.291502477
Client established: 2018-11-28 11:54:37.8387 +0800 CST m=+7.670990064
Client established: 2018-11-28 11:54:38.276112 +0800 CST m=+8.108388590
Client established: 2018-11-28 11:54:38.654549 +0800 CST m=+8.486814748
Client established: 2018-11-28 11:54:39.067531 +0800 CST m=+8.899784256
Client established: 2018-11-28 11:54:39.510984 +0800 CST m=+9.343223409
Client established: 2018-11-28 11:54:39.910885 +0800 CST m=+9.743112404
Client established: 2018-11-28 11:54:40.375044 +0800 CST m=+10.207258318
Client established: 2018-11-28 11:54:40.863303 +0800 CST m=+10.695502449
Client established: 2018-11-28 11:54:41.330891 +0800 CST m=+11.163076272
Client established: 2018-11-28 11:54:41.756931 +0800 CST m=+11.589104292
Client established: 2018-11-28 11:54:42.135702 +0800 CST m=+11.967863100
```
