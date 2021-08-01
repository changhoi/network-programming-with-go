package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ArgsCheck(2)

	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	CheckError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	CheckError(err)

	_, err = conn.Write([]byte("anything"))
	CheckError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	CheckError(err)

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
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
