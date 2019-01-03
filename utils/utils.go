package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"net"
)

func WritePkg(conn net.Conn, data []byte) (err error) {
	dataLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], dataLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("*********8")
		fmt.Println(err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("write data failed: ", err)
		return
	}
	return
}

func ReadPkg(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 8096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("读到的buf: ", buf[:4])

	pkgLen := binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		return
	}
	return
}
