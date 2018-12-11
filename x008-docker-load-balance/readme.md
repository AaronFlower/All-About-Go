## Docker Load-Balancing

1. 拉取镜像
```
docker pull nginx:1.15-alpine
```

2. 运行镜像
server 1
```
docker run --name nginx-test1 -d -p 8080:80 -v ~/go/src/github.com/aaronflower/ago/x008-docker-load-balance/home1:/usr/share/nginx/html nginx:1.15-alpine

```
server 2

```
docker run --name nginx-test2 -d -p 8081:80 -v ~/go/src/github.com/aaronflower/ago/x008-docker-load-balance/home2:/usr/share/nginx/html nginx:1.15-alpine
```

测试:
```
x008-docker-load-balance master ✗ 1h5m ◒ ➜ http :8080
HTTP/1.1 200 OK
Accept-Ranges: bytes
Connection: keep-alive
Content-Length: 224
Content-Type: text/html
Date: Tue, 11 Dec 2018 14:12:47 GMT
ETag: "5c0fc2c4-e0"
Last-Modified: Tue, 11 Dec 2018 13:59:32 GMT
Server: nginx/1.15.7

<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width" />
        <title>Home</title>
    </head>
    <body>
        <h1>Home 2</h1>
    </body>
</html>

x008-docker-load-balance master ✗ 1h6m ◒ ➜ http :8081
HTTP/1.1 200 OK
Accept-Ranges: bytes
Connection: keep-alive
Content-Length: 224
Content-Type: text/html
Date: Tue, 11 Dec 2018 14:12:50 GMT
ETag: "5c0fc2a6-e0"
Last-Modified: Tue, 11 Dec 2018 13:59:02 GMT
Server: nginx/1.15.7

<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width" />
        <title>Home</title>
    </head>
    <body>
        <h1>Home 1</h1>
    </body>
</html>
```

### 3. 布署 nginx 负载均衡
myweb.conf
```nginx
upstream.myupload.{¬
...server.localhost:8080.weight=5;¬
...server.localhost:8081.weight=1;¬
}¬
¬
server.{¬
....listen.8088;¬
¬
....server_name.localhost;¬
¬
....location./.{¬
........proxy_set_header.Host.$host;¬
........proxy_set_header.X-Real-IP.$remote_addr;¬
........proxy_set_header.X-Forwarded-For.$proxy_add_x_forwarded_for;¬
........root.../usr/share/nginx/html;¬
........index..index.html.index.htm;¬
........proxy_pass.http://myupload;¬
....}¬
}
```


### 4. 测试结果 
可以看出 weight 5; weight 1; 在请求 20 次的结果中, server 1 刚好是 server 2 的 5 倍。
```
 ➜ _ nginx -s reload
nginx ➜ bash -c 'i=0; while [ $i -lt 20 ]; do http :8088; i=$[$i + 1]; done;'|grep h1
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1>Home 2</h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1>Home 2</h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1>Home 2</h1>
        <h1> This is home 1 </h1>
        <h1> This is home 1 </h1>
        <h1>Home 2</h1>
```
