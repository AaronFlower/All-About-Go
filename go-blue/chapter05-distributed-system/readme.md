## Polls

### How to use MongoDB

```
# start the mongodb
$ mongod --dbpath ./db

# use mongodb cli
$ mongo
# or
$ mongo localhost
$ > show dbs
$ > help()
$ > db.help()
$ > db.polls.find()
$ > db.polls.help()
$ > db.polls.find({}).help()
```

在 `mongo cli` 中多用 `help()` 函数就对了。

- [ ] BSON

### How to use NSQ

```
# start the nsqd
$ nsqd

# start the nsqlookupd
$ nsqlookupd
```
- [ ] NSQ, nsqd, nsqlookupd
