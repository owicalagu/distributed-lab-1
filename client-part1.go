package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func connect() {
	conn, err := net.Dial("tcp", "3.86.143.115:34567")
	if err != nil {
		fmt.Println("error in connection")
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	_, _ = fmt.Fprintf(conn, text)

	for {
		fmt.Println("in for")
		text, _ := reader.ReadString('\n')
		_, _ = fmt.Fprintf(conn, text)
	}
	//_, _ = fmt.Fprintf(conn, "hi")

}

func main() {
	connect()
}
