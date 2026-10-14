package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	connStr := "host=localhost port=5432 user=zuhro password=qwerty dbname=coursedb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Cannot make pull:", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal("Connection failed:", err)
	}

	fmt.Println("Connection to data base succeded!")

	stats := db.Stats()
	fmt.Println("After .Ping():  ", stats.OpenConnections, "connections")
	fmt.Println("In use:", stats.InUse)
	fmt.Println("Idle:  ", stats.Idle)
	fmt.Println("Max connections:     ", stats.MaxOpenConnections)

}
