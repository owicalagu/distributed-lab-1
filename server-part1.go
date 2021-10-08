package main

import (
	"fmt"
	"net"
)

func newConnRead(listener net.Listener) {
	for {
		fmt.Println("in server for")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error in accepting")
		}
		fmt.Println()
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error in reading")
		}
		fmt.Println("got ", n, "bytes")
		fmt.Println(string(buf))
	}
}

func read() {
	for {

		listener, err := net.Listen("tcp", ":34567")
		fmt.Println("listening...")
		if err != nil {
			fmt.Println("error in listening")
			return
		}
		go newConnRead(listener)
	}

	/*for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error in accepting")
		}
		fmt.Println()
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error in reading")
		}
		fmt.Println("got ", n, "bytes")
		fmt.Println(string(buf))
	}*/
}

func main() {
	read()

}
