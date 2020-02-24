package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":7029")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	conn, err := ln.Accept()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
		fmt.Print("Received from client: ", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
