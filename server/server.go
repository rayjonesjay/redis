package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func Server() {

	port := ":6379"
	fmt.Printf("listening on port %s\n", strings.TrimPrefix(port, ":"))

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close() // close the connection once finished

	// create an infinite loop and receive commands from clients and responds to them
	for {

		// hold the message received, max size is 1024
		buf := make([]byte, 1024)

		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1) // fix this part, server should not break
		}

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}
