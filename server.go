package main

import (
	"fmt"
	"net"
	"os"

	db "github.com/ibenefit/db/sqlc"
)

const (
	SER_IP = "127.0.0.1"
	TYPE   = "tcp"
)

func serverStart(store db.Store, port string) {
	l, err := net.Listen(TYPE, SER_IP+":"+port)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("\nListening on " + SER_IP + ":" + port)
	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(store, conn)
	}
}

func handleRequest(store db.Store, conn net.Conn) {

	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// handle request here

	conn.Write([]byte("Message received."))

	conn.Close()
}

func send_request(port, req string) {

}
