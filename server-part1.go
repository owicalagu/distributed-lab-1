package main

import (
	"fmt"
	"net"
)

func read() {
	listener, err := net.Listen("tcp", ":22")
	if err != nil {
		fmt.Println("error in listening")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error in accepting")
		}
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error in reading")
		}
		fmt.Println("got ", n, "bytes")
		fmt.Println(string(buf))
	}
}

func main() {
	read()

}
