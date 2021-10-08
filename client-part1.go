package main

import (
	"fmt"
	"net"
)

func connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:8030")
	if err != nil {
		fmt.Println("error in connection")
		return
	}
	_, _ = fmt.Fprintf(conn, "hi")

}

func main() {
	connect()
}
