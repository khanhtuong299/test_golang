package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	db "github.com/ibenefit/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	db_src    = "postgresql://root:secret@localhost:5432/ibenefit?sslmode=disable"
	db_driver = "postgres"
)

func main() {

	conn, err := sql.Open(db_driver, db_src)
	if err != nil {
		log.Fatal("\ncannot connect to db:", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter port to create server (ex 9090): ")
	sport, _ := reader.ReadString('\n')
	fmt.Println("")
	store := db.NewStore(conn)

	go serverStart(store, sport)

	for {
		reader = bufio.NewReader(os.Stdin)
		fmt.Println("Enter port to connect server (ex): 9090")
		port, _ := reader.ReadString('\n')

		if sport == port {
			fmt.Println("\ncan not use the same port with server")
			continue
		}

		reader = bufio.NewReader(os.Stdin)
		fmt.Println("Enter request content")
		req, _ := reader.ReadString('\n')

		go send_request(port, req)
	}
}
