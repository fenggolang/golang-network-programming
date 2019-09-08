package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) { // 处理连接函数
	defer conn.Close()
	for {
		buf:=make([]byte,100)
		n,err:=conn.Read(buf) // 读取客户端数据
		if err!=nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("读取数据大小: %d,数据内容: %s\n",n,string(buf[0:n]))
		msg:=[]byte("hello,world\n")
		conn.Write(msg) // 发送数据
	}
}
func main() {
	fmt.Println("开始服务...")
	listener,err:=net.Listen("tcp","0.0.0.0:3000") // 1.创建监听
	if err!=nil{
		fmt.Println("监听失败!,消息是:",err)
		return
	}
	for {
		conn,err:=listener.Accept() // 2.接收客户端连接
		if err!=nil{
			fmt.Println("接收客户端连接失败")
			continue
		}
		go handle(conn) // 3.每一个客户端请求创建独立的协程(goroutine)来处理连接
	}
}
