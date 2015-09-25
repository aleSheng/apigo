# anlint的API server in Golang. 基于beego-demo

A web demo using Beego framework, with MongoDB,Redis support.

这是一个基于 [Beego](http://beego.me) 框架构建的应用 demo，后台数据库使用 [MongoDB](http://www.mongodb.org) 并使用 [Redis](http://redis.io) 存储 session 和一些统计数据。

## API列表

### 第一部分

该部分使用的数据库是 MongoDB 和 Redis。

在 static/test 目录下有如下的测试表单，除了用于测试外，也可看出具体的数据通讯协议：
* register.html
* login.html
* logout.html
* passwd.html
* uploads.html

说明：
* 输入数据通过 form 表单提交，返回数据均为 json。
* 使用 Beego 的 ParseForm 功能将输入数据解析到 struct 中。
* 使用 Beego 的 Validation 功能对数据进行校验。
* 使用 [scrypt](https://godoc.org/golang.org/x/crypto/scrypt) 算法进行密码处理。
* 由于 Beego 本身不支持多文件上传，故单独实现了 uploads API 来展示该功能，该功能与数据库无关。


说明：
* 参考 RESTful 模式设计 API。
* 输入数据采用 json，返回数据也是 json。
* 数据库操作使用原生 SQL，没有采用 ORM。


## 环境

### GO语言

包括安装 go，设置 $GOPATH 等，具体可参考：[How to Write Go Code](http://golang.org/doc/code.html)。

### MongoDB

在 conf/app.conf 中设置 MongoDB 参数，如：

```
[mongodb]
url = mongodb://127.0.0.1:27017/beego-demo
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

当前版本：

```
$ bee version
bee   :1.3.0
beego :1.5.0
Go    :go version go1.4.2 linux/amd64
```

## 运行

将代码放在 $GOPATH/src 目录下，运行（调试模式）：

```
$ cd $GOPATH/src/beego-demo/
$ bee run
```

正式部署时，可通过系统的 Init 服务来启动。在 scripts 目录下有 upstart 和 systemd 两套简易示例脚本，可参考使用。

例如，在 CentOS 6 下，复制 upstart/bdemo.conf 到 /etc/init/，相应修改后，执行：

```
# start bdemo
```

在 CentOS 7 下，复制 systemd/bdemo.service 到 /etc/systemd/system/，相应修改后，执行：

```
# systemctl daemon-reload
# systemctl enable bdemo.service
# systemctl start bdemo.service
```

由于 Init 是由 root 控制的，相应的服务缺省也具有 root 权限，故一般都应该做降权处理。可在 systemd 和 upstart 脚本中设置运行时的普通用户名和组名，具体可参考官方文档。

降权的问题在于普通用户无法绑定特权端口（如 80 ），不过实际环境下，还是建议在前面部署 Nginx 等成熟的 web 服务器，通过反向代理来访问应用。
