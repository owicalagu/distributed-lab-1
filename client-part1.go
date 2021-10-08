package main

import (
	"fmt"
	"net"
)

func connect() {
	conn, err := net.Dial("tcp", "172.31.95.134:3080")
	if err != nil {
		fmt.Println("error in connection")
		return
	}
	_, _ = fmt.Fprintf(conn, "hi")

}

func main() {
	connect()
}
