package network

import (
	"fmt"
	"net"
)

func Goruns() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		fmt.Println("Error listening", err.Error()+"\n")
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error()+"\n")
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error()+"\n")
			return //终止程序
		}
		fmt.Printf("Received data: %v \n", string(buf[:len]))
	}
}
