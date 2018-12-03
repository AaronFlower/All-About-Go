## Example 02
Use honey framework to implements a simple blog system.
Just for demo.

### Simple Test

```bash
△ ◒ ➜ http :9090
HTTP/1.1 200 OK
Content-Length: 11
Content-Type: text/plain; charset=utf-8
Date: Sat, 01 Dec 2018 13:41:44 GMT

Hello Again

 △ ◒ ➜ http POST :9090
HTTP/1.1 200 OK
Content-Length: 27
Content-Type: text/plain; charset=utf-8
Date: Sat, 01 Dec 2018 13:41:50 GMT

Your post has been handled.

△ ◒ ➜ http POST :9090/user
HTTP/1.1 200 OK
Content-Length: 14
Content-Type: text/plain; charset=utf-8
Date: Sat, 01 Dec 2018 13:41:56 GMT

Create a user!

△ ◒ ➜ http GET :9090/user
HTTP/1.1 200 OK
Content-Length: 17
Content-Type: text/plain; charset=utf-8
Date: Sat, 01 Dec 2018 13:42:02 GMT

Return all users!
```

### TODO
- Add [sqlx](https://github.com/jmoiron/sqlx)
