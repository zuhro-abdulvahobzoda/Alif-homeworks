package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Price  int
	Pages  int
}

func main() {
	// 1. Connect to PostgreSQL
	connStr := "host=localhost port=5432 user=zuhro password=qwerty dbname=coursedb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

	// 2. Add a new book using Exec() with all columns
	_, err = db.Exec("INSERT INTO book (title, author, price, pages) VALUES ($1, $2, $3, $4)", "Learning Go", "John Doe", 19, 320)
	if err != nil {
		log.Fatalf("Error inserting book: %v", err)
	}
	fmt.Println("New book added successfully using Exec().")

	// 3. Update the book's price using Exec()
	_, err = db.Exec("UPDATE book SET price = $1 WHERE title = $2", 24, "Learning Go")
	if err != nil {
		log.Fatalf("Error updating book price: %v", err)
	}
	fmt.Println("Book price updated successfully using Exec().")

	// 4. Delete a book by id using Exec()
	_, err = db.Exec("DELETE FROM book WHERE id = $1", 1)
	if err != nil {
		log.Fatalf("Error deleting book: %v", err)
	}
	fmt.Println("Book deleted successfully using Exec().")

	// 5. Get the id of a new book using INSERT ... RETURNING id with all columns
	var newBookID int
	err = db.QueryRow("INSERT INTO book (title, author, price, pages) VALUES ($1, $2, $3, $4) RETURNING id", "Advanced Go", "Jane Smith", 39, 450).Scan(&newBookID)
	if err != nil {
		log.Fatalf("Error inserting book and returning id: %v", err)
	}
	fmt.Printf("New book inserted with ID: %d\n", newBookID)

	// 6. Print all books to the console (selecting all 5 columns)
	rows, err := db.Query("SELECT id, title, author, price, pages FROM book")
	if err != nil {
		log.Fatalf("Error querying books: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n--- List of All Books ---")
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Price, &b.Pages); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}
		fmt.Printf("ID: %d | Title: %s | Author: %s | Price: %d | Pages: %d\n", b.ID, b.Title, b.Author, b.Price, b.Pages)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error during rows iteration: %v", err)
	}
}
