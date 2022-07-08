package main

import (
	"database/sql"
	"github.com/hanginfty/simple-bank/api"
	"log"
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgresql://root:passwd@localhost:5432/simple-bank?sslmode=disable"
	serverAddr = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to postgres db", err)
	}

	//api.NewServer(DB).Start(":8080")
	store := api.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
