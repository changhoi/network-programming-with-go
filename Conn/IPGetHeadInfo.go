package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ArgsCheck(2)

	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	CheckError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	CheckError(err)

	result, err := readFully(conn)
	CheckError(err)
	fmt.Println(string(result))

	os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}

func ArgsCheck(length int) {
	if len(os.Args) != length {
		fmt.Fprintf(os.Stderr, "Usage is wrong")
		os.Exit(1)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
