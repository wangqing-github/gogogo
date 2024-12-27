package net

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func StartServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error starting TCP server:", err.Error())
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Server listening on port 8081")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		// 使用 goroutine 处理并发连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	fmt.Println("有连接进来了")
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
		encode, err := Encode(msg)
		conn.Write(encode)
	}
}
