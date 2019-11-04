### haproxy介绍
```markdown
haproxy工作在前端用户和后端的Server之间，作为"中间人",haproxy会建立两个连接，一是用户端与haproxy建立
一个连接，另一个是haproxy与后端的server建立一个连接。

所有的proxy代理类服务的程序都会有一个相同的问题，就是出于proxy后端的server上不能够看到用户源IP地址，而
只能看到haproxy的IP地址。

haproxy作者Willy Tarreau开发了一个"Proxy protocol"用来向后端Server传递用户源IP地址。
目前Proxy protocol有v1和v2两个版本，v1偏重人类可读，v2是二进制格式，易于程序处理。
```