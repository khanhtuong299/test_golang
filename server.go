package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	_ "github.com/golang/protobuf/proto"
	db "github.com/ibenefit/db/sqlc"
	_ "github.com/ibenefit/utils"
	"google.golang.org/protobuf/proto"
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

	var data string

	raw_data := make([]byte, 1024)

	_, err := conn.Read(raw_data)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	err = proto.Unmarshal(raw_data, &data)

	if err != nil {
		fmt.Println("unmarshaling error: ", err)
		return
	}

	if validate_req(data) == false {
		fmt.Println(data)
		return
	} else {

	}

	conn.Write([]byte("Message received."))

	conn.Close()
}

func send_request(port, req string) {

	if validate_req(req) == false {
		fmt.Println("err: invalid request")
		return
	}

	conn, err := net.Dial(TYPE, SER_IP+":"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := proto.Marshal(req)
	if err != nil {
		fmt.Println("marshaling error: ", err)
		return
	}

	conn.Write([]byte(data))
}

func validate_req(req string) bool {

	s := strings.Split(req, " ")

	if len(s) != 3 {
		return false
	}

	if amount, err := strconv.Atoi(s[2]); err != nil || amount < 0 {
		return false
	}

	return true
}
