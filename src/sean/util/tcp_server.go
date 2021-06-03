package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			break
		}

		// 起一个协程去处理每个新连接
		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	defer conn.Close()
	packet := make([]byte, 1024)
	// 1. 读数据，没有可读数据，则阻塞
	_, _ = conn.Read(packet)

	// 2. 处理请求...
	s := string(packet)
	fmt.Printf("receive tcp connected: %s\n", s)

	// 3. 写返回，不可写则阻塞
	_, _ = conn.Write(packet)
}
