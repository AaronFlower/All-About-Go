## Nginx Reverse Proxy

Proxying 主要用于在多台服务器之间分发负载。用户可以无感知的访问网站内容。也可转发请求到应用服务器。

### Passing a Request to a Proxied Server
Nginx 代理接收到一个请求后，它会把这个请求发送到指定的服务器上，然后取回响应信息，再发回给客户。
Nginx 可以把这个请求转发到另一个 HTTP server 也可以是一个 Non-HTTP Server. 支持的协议有 FastCGI,
uwsgi, SCGI, memcached.

转发请求到一个 HTTP server, 要在 `location` 中使用 `proxy_pass` 命令。

```nginx
location /some/path {
    proxy_pass http://www.example.com/link/;
}

location ~ \.php {
    proxy_pass http://127.0.0.1:8000
}
```

在第一个例子中，被代理的服务器使用了 URI 的 PATH `/link/`, 那么 Nginx 会替换转发的请求。如：
`/some/path/page.html` 会被替换成 `http://www.example.com/link/page.html`.

代理一个 non-HTTP 服务器，只需要使用对应的 `**_pass` 命令就可以了. 支持的命令有:

- fastcgi_pass 转发请求到 FastCGI 服务器。
- uwsgi_pass
- scgi_pass
- memcached_pass

`proxy_pass` 命令可以指向一个组。可以用来做负载均衡。


### Passing Request Headers

可以使用 `proxy_set_header` 命令来转发 header.

### 配置 Buffers
命令：
1. proxy_buffering
2. proxy_buffer_size

### Choosing an Outgoing IP Address

有时候你可能需要根请求的相应 IP 来转发到相应的服务器上。你可以配置相应的 IP 网络或者 IP 地址段。
通过 `proxy_bind` 命令可以完成这项配置。
如：
```
location /app1/ {
    proxy_bind 127.0.0.1;
    proxy_pass http://example.com/app1/;
}

location /app2/ {
    proxy_bind 127.0.0.2;
    proxy_pass http://example.com/app2/;
}
```

```
location /app3/ {
    proxy_bind $server_addr;
    proxy_pass http://example.com/app3/;
}
```
