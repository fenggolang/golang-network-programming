package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:3000")
	if err!=nil{
		fmt.Println("拨号连接服务器失败:",err.Error())
		return
	}
	defer conn.Close()
	inputReader:=bufio.NewReader(os.Stdin)
	for {
	    fmt.Println("请输入...")
		str,_:=inputReader.ReadString('\n')
		data:=strings.Trim(str,"\n")
		if data == "quit" { // 输入quit退出
			return
		}
		_,err:=conn.Write([]byte(data)) // 发送数据
		if err!=nil{
			fmt.Println("发送数据失败:",err)
			return
		}
		buf:=make([]byte,512)
		n,err:=conn.Read(buf) // 读取服务器端数据
		fmt.Println("读取到的服务器端的数据是:",string(buf[:n]))
	}
}
