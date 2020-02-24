package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Usage:\n$ go-telnet [host] [port]\n")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	_, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

}
