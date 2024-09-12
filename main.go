package main

import (
	"fmt"
	"net"
	"io"
)

func main() {
	fmt.Println("Listening on port :6379")
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Printf("Error listening on port :6379, err= %s\n", err)
		return
	}

	//Listen for incoming connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Printf("Error accepting connection: %s\n", err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		// read message from client
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				return
			} else {
				fmt.Printf("Error reading from client: %s\n", err)
				continue
			}
		}

		// write message to client
		conn.Write([]byte("+PONG\r\n"))
	}
}