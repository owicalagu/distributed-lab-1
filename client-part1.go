package main

import (
	"fmt"
	"net"
)

func connect() {
	conn, err := net.Dial("tcp", "172.31.95.134:22")
	if err != nil {
		fmt.Println("error in connection")
		return
	}
	_, _ = fmt.Fprintf(conn, "hi")

}

func main() {
	connect()
}
