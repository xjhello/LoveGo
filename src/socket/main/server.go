package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("recv error : %v\n", err)
		}
		revc := string(buf[:n])
		fmt.Printf("recv data: %v \n", revc)
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("write from conn failed, err:%v\n", err)
			break
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Printf("socket error!")
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("acc")
			continue
		}
		go process(conn)
	}
}
