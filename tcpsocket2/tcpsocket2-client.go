package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"golang-network-programming/tcpsocket2/config"
)

// 首先通过Dial建立与服务器的连接，之后读取标标准输入的行，将其传递给服
// 务器，然后从服务器读取响应。这里处理粘包采用了自己封包的方式。
func main() {

	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > config.SERVER_RECV_LEN {
				toWrite = line[written : written+config.SERVER_RECV_LEN]
			} else {
				toWrite = line[written:]
			}
			n, err = conn.Write([]byte(toWrite))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("请求:", toWrite)
			msg := make([]byte, config.SERVER_RECV_LEN)
			n, err = conn.Read(msg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println("响应:", string(msg))
		}
	}
}
