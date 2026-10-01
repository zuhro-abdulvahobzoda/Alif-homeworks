package main

import (
	in "fmt"
	b "library/book"
	c "library/catalog"
)

func main() {
	books := []b.Book{
		{
			Title:  "The Hobbit",
			Author: "J.R.R. Tolkien",
			Year:   1937,
		},
		{
			Title:  "To Kill a Mockingbird",
			Author: "Harper Lee",
			Year:   1960,
		},
		{
			Title:  "1984",
			Author: "George Orwell",
			Year:   1949,
		},
		{
			Title:  "Pride and Prejudice",
			Author: "Jane Austen",
			Year:   1813,
		},
		{
			Title:  "The Catcher in the Rye",
			Author: "J.D. Salinger",
			Year:   1951,
		},
		{
			Title:  "Fahrenheit 451",
			Author: "Ray Bradbury",
			Year:   1953,
		},
		{
			Title:  "Brave New World",
			Author: "Aldous Huxley",
			Year:   1932,
		},
		{
			Title:  "Moby-Dick",
			Author: "Herman Melville",
			Year:   1851,
		},
		{
			Title:  "The Lord of the Rings",
			Author: "J.R.R. Tolkien",
			Year:   1954,
		},
		{
			Title:  "Animal Farm",
			Author: "George Orwell",
			Year:   1945,
		},
		{
			Title:  "Frankenstein",
			Author: "Mary Shelley",
			Year:   1818,
		},
		{
			Title:  "The Picture of Dorian Gray",
			Author: "Oscar Wilde",
			Year:   1890,
		},
		{
			Title:  "One Hundred Years of Solitude",
			Author: "Gabriel García Márquez",
			Year:   1967,
		},
		{
			Title:  "Crime and Punishment",
			Author: "Fyodor Dostoevsky",
			Year:   1866,
		},
	}

	in.Printf("Seaching %s's book(s)...\n", "J.R.R Tolkien")
	in.Println(c.FindByAuthor(books, "J.R.R. Tolkien"))
}
