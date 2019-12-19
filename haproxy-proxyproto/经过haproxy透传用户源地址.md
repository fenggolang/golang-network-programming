### haproxy介绍
```markdown
haproxy工作在前端用户和后端的Server之间，作为"中间人",haproxy会建立两个连接，一是用户端与haproxy建立
一个连接，另一个是haproxy与后端的server建立一个连接。

proxy protocol是比较新的协议，但目前已经有很多软件支持，如haproxy、nginx、apache、squid、mysql等等，要使用
proxy protocol需要两个角色sender和receiver,sender在与receiver之间建立连接后，会先发送一个带有客户信息的tcp header,
因为更改了tcp协议，需receiver也支持proxy protocol，否则不能识别tcp包头，导致无法成功建立连接。

所有的proxy代理类服务的程序都会有一个相同的问题，就是处于proxy后端的server上不能够看到用户源IP地址，而
只能看到haproxy的IP地址。

haproxy作者Willy Tarreau开发了一个"Proxy protocol"用来向后端Server传递用户源IP地址。
目前Proxy protocol有v1和v2两个版本，v1偏重人类可读，v2是二进制格式，易于程序处理。
```
[proxy protocol - 让haproxy反向代理的后端nginx拿到用户真实ip](https://wangbin.io/blog/it/haproxy-proxy-protocol.html)
### 测试haproxy proxyproto协议
```markdown
# 1.写一个listen监听器server,使用proxyproto.Listener包装，然后把监听器程序放到paas上部署，对外暴露
7层路由和4层路由，
# 2.修改haproxy模板配置文件，如果四层中的annotation注解中有如下,haproxy模板配置文件会读取其注解的值，写入对应四层路由的server后面:
## 4层路由注解
    m8cloud/proxyproto-receiver: ''
    m8cloud/proxyproto-sender: send-proxy
## haproxy模板配置获取注解的值
  bind 0.0.0.0:{{index $cfg.Annotations "m8cloud/port"}} {{index $cfg.Annotations "m8cloud/proxyproto-receiver"}}
  ...
  server {{$endpoint.ID}} {{$endpoint.IP}}:{{$endpoint.Port}} weight {{$weight}} {{ if (eq (index $cfg.Annotations "m8cloud/proxyproto-sender") "send-proxy")}}send-proxy

不使用haproxy,直接在服务器上部署这个监听器程序，然后本机电脑访问，发现打印的的地址就是源ip地址和源端口，正确!
```
### ping和telnet
```markdown
ping telnet的区别
　ping: 用来检查网络是否通畅或网络连接速度(Ping域名可以得出解析IP)
　telnet: 用来检查指定ip是否开放指定端口的
说明：
Ping不通并不一定代表网络不通。ping是基于ICMP协议的命令，就是你发出去一个数据包，对方收到后返给你一个！就好比声纳。这个协议是可以禁止的！禁止后，如果你ping对方，对方收到后就不回馈给你，这样你就显示无法ping通，但实际你们还是连着的！telnet是登陆服务器的！服务没禁止就能登陆

clientIP clientPort
backendIP backendPort
proxyIP proxyPort
serverIP serverPort
```

### 测试源代码地址

```bash
https://github.com/fenggolang/golang-network-programming.git
```

