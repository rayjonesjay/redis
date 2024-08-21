package server

import (
	"fmt"
	"net"
	"redis/resp"
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

		resp := resp.NewResp(conn)

		value, err := resp.Read()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(value)

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}
