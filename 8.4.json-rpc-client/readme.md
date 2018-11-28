## HTTP RPC
### Client Test
```
bash -c 'i=0; while [ $i -lt 10 ]; do go run main.go :1234 $[$i + 2] $i; i=$[$i + 1]; done;'
Arith: 2 * 0 = 0
2018/11/28 17:33:17 arith error:divide by zero
exit status 1
Arith: 3 * 1 = 3
Arith: 3 / 1 = 3 remainder 0
Arith: 4 * 2 = 8
Arith: 4 / 2 = 2 remainder 0
Arith: 5 * 3 = 15
Arith: 5 / 3 = 1 remainder 2
Arith: 6 * 4 = 24
Arith: 6 / 4 = 1 remainder 2
Arith: 7 * 5 = 35
Arith: 7 / 5 = 1 remainder 2
Arith: 8 * 6 = 48
Arith: 8 / 6 = 1 remainder 2
Arith: 9 * 7 = 63
Arith: 9 / 7 = 1 remainder 2
Arith: 10 * 8 = 80
Arith: 10 / 8 = 1 remainder 2
Arith: 11 * 9 = 99
Arith: 11 / 9 = 1 remainder 2
```
### Server Output
```
TCP RPC has been listening at :1234
Client Established:  2018-11-28 17:33:17.278438 +0800 CST m=+3.002656614
Client Established:  2018-11-28 17:33:18.054217 +0800 CST m=+3.778435061
Client Established:  2018-11-28 17:33:18.892704 +0800 CST m=+4.616921961
Client Established:  2018-11-28 17:33:19.72151 +0800 CST m=+5.445728241
Client Established:  2018-11-28 17:33:20.495824 +0800 CST m=+6.220042173
Client Established:  2018-11-28 17:33:21.277076 +0800 CST m=+7.001294488
Client Established:  2018-11-28 17:33:22.001679 +0800 CST m=+7.725897158
Client Established:  2018-11-28 17:33:22.77848 +0800 CST m=+8.502698565
Client Established:  2018-11-28 17:33:23.665 +0800 CST m=+9.389217990
Client Established:  2018-11-28 17:33:24.450549 +0800 CST m=+10.174767240
```
