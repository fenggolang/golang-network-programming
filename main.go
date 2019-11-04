package main

import (
	"fmt"
	"net"

	"github.com/armon/go-proxyproto"
)

func main() {
	// Create a listener
	list, err := net.Listen("tcp", "...")
	if err != nil {
		fmt.Printf("Listen失败:%v", err)
		return
	}
	// Wrap listener in a proxyproto listener
	proxyList := &proxyproto.Listener{Listener: list}
	conn, err := proxyList.Accept()
	if err != nil {
		fmt.Printf("Accept失败:%v", err)
		return
	}
	addr := conn.RemoteAddr()
	fmt.Println("addr", addr.Network())
	fmt.Println("addr", addr.String())

}
