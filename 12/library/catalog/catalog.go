package catalog

import (
	b "library/book"
)

func FindByAuthor(books []b.Book, author string) []b.Book {
	foundB := []b.Book{}
	for _, bk := range books {
		if author == bk.Author {
			foundB = append(foundB, bk)

		}

	}
	return foundB
}
