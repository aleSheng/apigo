# anlint的API server in Golang. 基于beego-demo

A web demo using Beego framework, with MongoDB,Redis support.

这是一个基于 [Beego](http://beego.me) 框架构建的应用 demo，后台数据库使用 [MongoDB](http://www.mongodb.org) 并使用 [Redis](http://redis.io) 存储 session 和一些统计数据。

## API列表

API都在routers/router.go里面。
V层：该server纯作为API server（json），不用html template render。
C层：在controllers目录里面是数据处理和业务逻辑。
M层：在models目录下有数据描述。


### 数据库

该部分使用的数据库是 MongoDB 和 Redis。
mongodb：数据持久化
redis：session，以及数据分析

说明：
* 使用 Beego 的 ParseForm 功能将输入数据解析到 struct 中。
* 使用 Beego 的 Validation 功能对数据进行校验。
* 使用 [scrypt](https://godoc.org/golang.org/x/crypto/scrypt) 算法进行密码处理。
* 由于 Beego 本身不支持多文件上传，故单独实现了 uploads API 来展示该功能，该功能与数据库无关。


说明：
* 参考 RESTful 模式设计 API。
* 输入数据采用 json，返回数据也是 json。


## 环境

### GO语言

包括安装 go，设置 $GOPATH 等，具体可参考：[How to Write Go Code](http://golang.org/doc/code.html)。

### MongoDB

在 conf/app.conf 中设置 MongoDB 参数，如：

```
[mongodb]
url = mongodb://127.0.0.1:27017/anlintdb
```

完整的 url 写法可参考：http://godoc.org/gopkg.in/mgo.v2#Dial

这里单独封装了一个 mymongo 包来实现数据库的初始化，以简化后续的数据库操作。


### Redis

在 conf/app.conf 中设置 Redis 参数，涉及两个地方，一个是 session，一个是 cache，两者可以不同：

```
sessionsavepath = 127.0.0.1:6379

[cache]
server = 127.0.0.1:6379
password =
```
这里单独封装了一个 myredis 包来实现数据库的初始化，以简化后续的数据库操作。

### Beego

安装/升级所有依赖包：

```
$ go get -u github.com/astaxie/beego
$ go get -u github.com/beego/bee
$ go get -u github.com/astaxie/beego/session/redis
$ go get -u gopkg.in/mgo.v2
$ go get -u github.com/garyburd/redigo/redis
$ go get -u golang.org/x/crypto/scrypt
```


## 运行

将代码放在 $GOPATH/src 目录下，运行（调试模式）

｀｀｀
bee run
｀｀｀

正式部署时，可用supervisor守护
