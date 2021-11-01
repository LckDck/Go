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
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	for {
		//Some changes here
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			return
		}
		fmt.Fprintf(conn, text+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			return
		}
		fmt.Fprintln(os.Stdout, "Answer from server: "+message)
	}
}
