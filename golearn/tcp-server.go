package golearn

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func TcpConnect() {

	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {

	defer conn.Close()

	for {

		var buf [128]byte

		reader := bufio.NewReader(conn)

		lenBytes, err := reader.Peek(4)

		buffer := bytes.NewBuffer(lenBytes)

		var length int32
		binary.Read(buffer, binary.LittleEndian, &length)

		

		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		fmt.Println("接收到数据：%s ", string(buf[:n]))
	}

}
