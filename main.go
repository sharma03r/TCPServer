package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err.Error())
	}
	fmt.Println("Hello, World! ", conn)
	defer conn.Close()
}
