package tcp

import (
	"encoding/binary"
	"fmt"
	"net"
)

func Dial() {

	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	defer conn.Close()

	for {
		var s string
		fmt.Scan(&s)
		l := len(s)
		b := make([]byte, 0, l+4)

		binary.LittleEndian.PutUint32(b, uint32(l))
		b = append(b, []byte(s)...)

		n, err := conn.Write(b)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Printf("成功发送%d字节", n)
		}
	}

}
