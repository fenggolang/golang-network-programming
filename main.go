package main

import (
	"fmt"
	"net"

	"github.com/armon/go-proxyproto"
)

/**
服务端
 */
func main() {
	// 创建一个监听器
	//listener, err := net.Listen("tcp", "172.18.1.242:9907")
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
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
		fmt.Printf("经过haproxy透传用户源地址IP=%s\n", addr.IP.String())
		fmt.Printf("经过haproxy透传用户源端口PORT=%d\n", addr.Port)
		conn.Close()
	}
}
