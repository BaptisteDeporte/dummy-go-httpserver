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
	const bufferSize = 1024;
	
    defer connection.Close()

    fmt.Println(connection.RemoteAddr())

    buffer := make([]byte, bufferSize)
    var data bytes.Buffer

    for {
        n, err := connection.Read(buffer)

        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Erreur de lecture :", err)
            break
        }

        fmt.Println("Nombre d'octets lus :", n)
        data.Write(buffer[:n])

		if n < bufferSize {
			break;
		}
    }

    fmt.Print("Données lues : ", data.String())

    message := []byte("HTTP/1.1 200 OK\r\nContent-Length:2\r\n\r\nOK")
    _, err := connection.Write(message)
    if err != nil {
        fmt.Println("Erreur d'écriture :", err)
    }
}
