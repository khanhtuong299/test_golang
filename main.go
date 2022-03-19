package main

import (
	"bufio"
	"context"
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
	store := db.NewStore(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Create 3 account A1 A2 A3 (only yes when start first server) (y/n): ")
	create_acc, _ := reader.ReadString('\n')

	if create_acc == "y" {
		create_acc_init(store)
	}

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter port to create server (ex 9090): ")
	sport, _ := reader.ReadString('\n')
	fmt.Println("")

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
		fmt.Println("request format: [account 1] [account 2] [amount]")
		fmt.Println("Ex: A1 A2 10")
		req, _ := reader.ReadString('\n')

		go send_request(port, req)
	}
}

func create_acc_init(store db.Store) {
	arg1 := db.CreateAccountParams{
		Account:    "A1",
		PublicKey:  "1234",
		PrivateKey: "4567",
	}

	arg2 := db.CreateAccountParams{
		Account:    "A2",
		PublicKey:  "1234",
		PrivateKey: "4567",
	}

	arg3 := db.CreateAccountParams{
		Account:    "A3",
		PublicKey:  "1234",
		PrivateKey: "4567",
	}

	store.CreateAccount(context.Background(), arg1)
	store.CreateAccount(context.Background(), arg2)
	store.CreateAccount(context.Background(), arg3)
}
