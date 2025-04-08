package main

import (
	"fmt"
	"net"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Listening incoming connections")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Println(connection.RemoteAddr())
	message := []byte("ok")
	connection.Write(message)
	err := connection.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}