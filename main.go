package main

import (
	"bytes"
	"fmt"
	"io"
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
	message := []byte("listening")
	connection.Write(message)

	buffer := make([]byte, 1024)
	var data bytes.Buffer

	for {
        n, err := connection.Read(buffer)
        if err != nil {
            if err == io.EOF {
                // Fin des données
                break
            }
            fmt.Println("Erreur de lecture :", err)
            return
        }
        data.Write(buffer[:n])
    }

    fmt.Println("input data :", data.String())
}