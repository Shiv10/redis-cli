package main

import (
	"fmt"
	"net"
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
		
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// write message to client
		conn.Write([]byte("+PONG\r\n"))
	}
}