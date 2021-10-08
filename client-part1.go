package main

import (
	"fmt"
	"net"
)

func connect() {
	conn, err := net.Dial("tcp", "172.31.95.134:8010")
	if err != nil {
		fmt.Println("error in connection")
		fmt.Println(err)
		return
	}
	_, _ = fmt.Fprintf(conn, "hi")

}

func main() {
	connect()
}
