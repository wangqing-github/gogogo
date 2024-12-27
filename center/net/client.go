package net

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func StartClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	sendMsg(conn)
}

// 模拟每5秒向服务器发送一次信息
func sendMsg(conn net.Conn) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
			msg := `Hello, Hello. How are you?`
			data, err := Encode(msg)
			if err != nil {
				fmt.Println("encode msg failed, err:", err)
				return
			}
			_, err = conn.Write(data)
			if err != nil {
				fmt.Println("Error writing to server:", err.Error())
				break
			}
			fmt.Println("向服务器发送信息:", data)
			// 读取服务器的回显
			reader := bufio.NewReader(conn)
			decode, err := Decode(reader)
			if err != nil {
				fmt.Println("Error reading from server:", err.Error())
			}
			fmt.Println("Server response:", decode)
		}
	}
}
