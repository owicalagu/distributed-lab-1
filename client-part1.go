package main

import (
	"fmt"
	"net"
)

func connect() {
	conn, err := net.Dial("tcp", "3.86.143.115:3080")
	if err != nil {
		fmt.Println("error in connection")
		return
	}
	fmt.Fprintf(conn, "hi")

}

func main() {

}
