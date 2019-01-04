package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"net"
)

type Transer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transer) WritePkg(data []byte) (err error) {
	dataLen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], dataLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("*********8")
		fmt.Println(err)
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("write data failed: ", err)
		return
	}
	return
}

func (this *Transer) ReadPkg() (msg message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("读到的this.Buf: ", this.Buf[:4])

	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	err = json.Unmarshal(this.Buf[:pkgLen], &msg)
	if err != nil {
		return
	}
	return
}
