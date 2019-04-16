## How to use

### Build The Echo Server

```
go build
```

### Run the Server

```
./01-uds
```

### Use `nc` to test the server
```
nc -U /tmp/echo.sock
```

### Check the `/tmp/echo.sock`

```
ll /tmp/echo.sock
srwxr-xr-x  1 aaron  wheel     0B Apr 16 11:36 /tmp/echo.sock
```
