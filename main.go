package main

import (
	"fmt"
	"net"

	"github.com/armon/go-proxyproto"
)

/**
测试程序，把程序放到paas上部署(paas上的haproxy的配置需要做四层负载注解部分修改)
*/
func main() {
	// 创建一个监听器
	listener, err := net.Listen("tcp4", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("Listen失败:%v", err)
		return
	}
	// 用proxyproto 监听器包装原生监听器
	proxyListener := &proxyproto.Listener{Listener: listener}
	for {
		conn, err := proxyListener.Accept()
		if err != nil {
			fmt.Printf("Accept失败:%v", err)
			return
		}
		addr := conn.RemoteAddr().(*net.TCPAddr)
		fmt.Printf("经过haproxy透传用户源地址=%s\n", addr.IP.String())
		fmt.Printf("经过haproxy透传用户源端口=%d\n", addr.Port)
		//resp:=fmt.Sprintf("经过haproxy透传用户源地址=%s,源端口=%d",addr.IP.String(),addr.Port)
		resp := fmt.Sprintf("clientIP=%s,clientPort=%d", addr.IP.String(), addr.Port)
		conn.Write([]byte(resp))
		conn.Close()
	}
}
